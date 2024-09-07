import { NextResponse } from 'next/server';

export async function middleware(req) {
  const url = req.nextUrl.pathname;
  const cookie = req.cookies.get('session_token')?.value;
  const baseUrl = req.nextUrl.origin;
  const apiUrl = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';

  console.log('Session Token:', cookie);

  // Check if the URL is accessing a specific group by ID

  // Existing session validation logic
  if (url === '/auth' || url === '/auth/login') {
    if (cookie) {
      try {
        const res = await fetch(`${apiUrl}/validatecookie`, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ cookie }),
        });

        const data = await res.json();

        if (data.valid) {
          const response = NextResponse.redirect(`${baseUrl}/`);
          response.cookies.set('currentUserID', data.currentUserID, { httpOnly: true });
          return response;
        } else {
          const response = NextResponse.redirect(`${baseUrl}/auth/login`);
          response.cookies.set('currentUserID', '', { httpOnly: true }); // Vider currentUserID
          return response;
        }
      } catch (error) {
        console.error('Error validating cookie:', error);
      }
    }
    return NextResponse.next();
  } else {
    if (!cookie) {
      return NextResponse.redirect(`${baseUrl}/auth/login`);
    }

    try {
      const res = await fetch(`${apiUrl}/validatecookie`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ cookie }),
      });

      const data = await res.json();

      if (!data.valid) {
        const response = NextResponse.redirect(`${baseUrl}/auth/login`);
        response.cookies.set('currentUserID', '', { httpOnly: true }); // Vider currentUserID
        return response;
      }

      if (url.startsWith('/groups/')) {
        const id = url.split('/groups/')[1];
        
        // Validate the ID format (UUID) before making a request
        const isValidUUID = /^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$/.test(id);
        
        if (!isValidUUID) {
          // Redirect to Not Found page if the ID format is not valid
          return NextResponse.redirect(`${baseUrl}/404`);
        }
    
        try {
          // Make a request to the backend to verify if the ID exists
          const res = await fetch(`${apiUrl}/group/getEvents?groupId=${id}`, {
            method: 'GET',
            credentials: 'include'
          });
          console.log("res status",res.status);
          
          if (res.status === 404) {
            // Redirect to Not Found page if the ID does not exist
            return NextResponse.redirect(`${baseUrl}/404`);
          }
    
          // Continue to the group page if the ID is valid
          return NextResponse.next();
        } catch (error) {
          console.error('Error verifying group ID:', error);
          // Redirect to Not Found page in case of any errors during verification
          return NextResponse.redirect(`${baseUrl}/404`);
        }
      }

      const response = NextResponse.next();
      response.cookies.set('currentUserID', data.currentUserID, { httpOnly: true });
      return response;
    } catch (error) {
      console.error('Erreur lors de la validation du cookie:', error);
      return NextResponse.redirect(`${baseUrl}/auth/login`);
    }
  }
}

export const config = {
  matcher: ['/((?!api|_next/static|_next/image|favicon.ico).*)'],
};