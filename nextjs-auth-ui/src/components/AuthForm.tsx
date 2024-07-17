"use client"; // This directive marks the component as a Client Component

import { useState } from 'react';
import { signup, login } from '../utils/auth';
import axios from 'axios';

interface AuthFormProps {
    mode: 'login' | 'signup';
    onOtpSent: () => void;
}

const AuthForm: React.FC<AuthFormProps> = ({ mode, onOtpSent }) => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState('');
    const [otp, setOtp] = useState(['', '', '', '', '', '']);
    const [isOtpSent, setIsOtpSent] = useState(false);

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        try {
            if (mode === 'login') {
                await login(username, password);
            } else {
                await signup(username, password);
                setIsOtpSent(true);
                onOtpSent();
            }
        } catch (err) {
            if (axios.isAxiosError(err)) {
                setError(err.response?.data || 'An error occurred');
            } else {
                setError('An error occurred');
            }
        }
    };

    const handleOtpChange = (index: number) => (e: React.ChangeEvent<HTMLInputElement>) => {
        const newOtp = [...otp];
        newOtp[index] = e.target.value;
        setOtp(newOtp);
    };

    const handleVerifyOtp = async () => {
        const otpCode = otp.join('');
        // Call the verification endpoint with the OTP
        try {
            // Replace with the actual verification logic
            await axios.post('/verify', { username, code: otpCode });
            // Handle successful verification
        } catch (err) {
            setError('Invalid OTP or verification failed');
        }
    };

    return (
        <form onSubmit={handleSubmit} className="space-y-4">
            {!isOtpSent ? (
                <>
                    <div>
                        <label htmlFor="username" className="block text-sm font-medium text-gray-700">
                            Email
                        </label>
                        <input
                            type="email"
                            id="username"
                            value={username}
                            onChange={(e) => setUsername(e.target.value)}
                            required
                            className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm text-black"
                        />
                    </div>
                    <div>
                        <label htmlFor="password" className="block text-sm font-medium text-gray-700">
                            Password
                        </label>
                        <input
                            type="password"
                            id="password"
                            value={password}
                            onChange={(e) => setPassword(e.target.value)}
                            required
                            className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm text-black"
                        />
                    </div>
                    {error && <p className="text-red-500 text-sm">{error}</p>}
                    <div>
                        <button
                            type="submit"
                            className="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
                        >
                            {mode === 'login' ? 'Log In' : 'Sign Up'}
                        </button>
                    </div>
                </>
            ) : (
                <>
                    <div className="flex space-x-2">
                        {otp.map((digit, index) => (
                            <input
                                key={index}
                                type="text"
                                value={digit}
                                onChange={handleOtpChange(index)}
                                maxLength={1}
                                className="w-10 px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm text-black text-center"
                            />
                        ))}
                    </div>
                    {error && <p className="text-red-500 text-sm">{error}</p>}
                    <div>
                        <button
                            type="button"
                            onClick={handleVerifyOtp}
                            className="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
                        >
                            Verify OTP
                        </button>
                    </div>
                </>
            )}
        </form>
    );
};

export default AuthForm;
