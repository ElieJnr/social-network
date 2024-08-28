import { NextResponse } from 'next/server';

export function middleware(req) {
  
  const token = req.cookies.get('session_token');

  if (!token) {
    const url = req.nextUrl.clone();
    return NextResponse.redirect(new URL("/auth", req.url));
  }

  console.log('Token valid, continuing to the page');
  return NextResponse.next();
}

// Export config sans matcher pour voir si le middleware est appelé
export const config = {
  matcher: ['/', '/profile'],
};
