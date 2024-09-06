import { NextResponse } from 'next/server';

export async function middleware(req) {
  const url = req.nextUrl.pathname;
  const cookie = req.cookies.get('session_token')?.value; // Utilisation de l'opérateur optionnel

  console.log('Session Token:', cookie);

  // Construisez l'URL absolue de redirection
  const baseUrl = req.nextUrl.origin;

  if (url === '/auth' || url === '/auth/login') {
    // Pour /auth : empêcher l'accès si le cookie est valide
    if (cookie) {
      try {
        const res = await fetch('http://localhost:8080/validatecookie', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ cookie }), // Assurez-vous d'envoyer le bon format de données
        });

        const data = await res.json();

        if (data.valid) {
          // Redirigez vers la page d'accueil si le cookie est valide
          return NextResponse.redirect(`${baseUrl}/`);
        }
      } catch (error) {
        console.error('Erreur lors de la validation du cookie:', error);
      }
    }
    // Si le cookie n'existe pas ou est invalide, continuer vers /auth
    return NextResponse.next();
  } else {
    // Pour les autres pages : vérifier la validité du cookie et rediriger si invalide
    if (!cookie) {
      return NextResponse.redirect(`${baseUrl}/auth/login`);
    }

    try {
      const res = await fetch('http://localhost:8080/validatecookie', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ cookie }), // Assurez-vous d'envoyer le bon format de données
      });

      const data = await res.json();
      if (!data.valid) {
        return NextResponse.redirect(`${baseUrl}/auth/login`);
      }

      // Si le cookie est valide, continuer vers la page demandée
      return NextResponse.next();
    } catch (error) {
      console.error('Erreur lors de la validation du cookie:', error);
      // En cas d'erreur, rediriger vers la page de login
      return NextResponse.redirect(`${baseUrl}/auth/login`);
    }
  }
}

export const config = {
  matcher: ['/((?!api|_next/static|_next/image|favicon.ico).*)'],
};

