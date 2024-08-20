import { NextResponse } from 'next/server';

export async function POST(req) {
  try {
    const body = await req.json();
    
    const response = await fetch('http://localhost:8080/users/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(body),
    });

    const data = await response.json();

    const cookies = response.headers.get('set-cookie');
    const res = NextResponse.json(data, { status: response.status });

    if (cookies) {
      res.headers.set('Set-Cookie', cookies);
    }

    return res;
  } catch (error) {
    return NextResponse.json(
      { error: 'Une erreur est survenue lors de la connexion' },
      { status: 500 }
    );
  }
}
