"use client"; // This directive marks the component as a Client Component

import { useState } from 'react';
import AuthForm from '../components/AuthForm';

export default function Home() {
  const [mode, setMode] = useState<'signup' | 'login'>('signup');

  return (
      <main className="flex min-h-screen flex-col items-center justify-center p-6 bg-gray-50">
        <div className="w-full max-w-md space-y-8">
          <div>
            <h2 className="mt-6 text-center text-3xl font-extrabold text-gray-900">
              {mode === 'signup' ? 'Sign Up' : 'Log In'}
            </h2>
          </div>
          <AuthForm mode={mode} />
          <div className="flex justify-center">
            <button
                type="button"
                className="mt-4 text-indigo-600 hover:text-indigo-900"
                onClick={() => setMode(mode === 'signup' ? 'login' : 'signup')}
            >
              {mode === 'signup' ? 'Already have an account? Log In' : "Don't have an account? Sign Up"}
            </button>
          </div>
        </div>
      </main>
  );
}
