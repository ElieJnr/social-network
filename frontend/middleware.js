import { NextResponse } from 'next/server';

export function middleware(req) {
  const token = req.cookies.get('session_token');
  const url = req.nextUrl.clone();
  
  if (token && (url.pathname === '/auth' || url.pathname === '/auth/login')) {
    return NextResponse.redirect(new URL('/', req.url));
  }

  if (!token && (url.pathname !== '/auth' && url.pathname !== '/auth/login')) {
    return NextResponse.redirect(new URL('/auth/login', req.url));
  }

  console.log('Token valid, continuing to the page');
  return NextResponse.next();
}

export const config = {
  matcher: ['/', '/profile', '/auth', '/auth/login'],
};
