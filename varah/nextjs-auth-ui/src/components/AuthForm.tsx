// src/components/AuthForm.tsx

"use client"; // This directive marks the component as a Client Component

import { useState, useRef } from 'react';
import { signup, login } from '../utils/auth';
import axios from 'axios';
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import GitHubLoginButton from './GitHubLoginButton';
import Header from './Header';

interface AuthFormProps {
    mode: 'login' | 'signup';
    onOtpSent: () => void;
}

const AuthForm: React.FC<AuthFormProps> = ({ mode, onOtpSent }) => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [otp, setOtp] = useState(['', '', '', '', '', '']);
    const [isOtpSent, setIsOtpSent] = useState(false);
    const otpRefs = useRef<(HTMLInputElement | null)[]>([]);

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        try {
            if (mode === 'login') {
                await login(username, password);
                toast.success('Logged in successfully!');
            } else {
                await signup(username, password);
                setIsOtpSent(true);
                toast.success('Verification email sent. Please check your inbox.');
                onOtpSent();
            }
        } catch (err) {
            if (axios.isAxiosError(err)) {
                toast.error(err.response?.data || 'An error occurred');
            } else {
                toast.error('An error occurred');
            }
        }
    };

    const handleOtpChange = (index: number) => (e: React.ChangeEvent<HTMLInputElement>) => {
        const newOtp = [...otp];
        newOtp[index] = e.target.value;
        setOtp(newOtp);

        // Auto tab to next input
        if (e.target.value.length === 1 && index < otpRefs.current.length - 1) {
            otpRefs.current[index + 1]?.focus();
        }
    };

    const handleVerifyOtp = async () => {
        const otpCode = otp.join('');
        console.log(`Verifying OTP: ${otpCode}`); // Debug log for OTP code
        // Call the verification endpoint with the OTP
        try {
            const response = await axios.post('http://localhost:8080/verify', { username, code: otpCode });
            console.log(`Verification response: ${response.data}`); // Debug log for response
            // Handle successful verification
            toast.success('OTP verified successfully!');
        } catch (err) {
            if (axios.isAxiosError(err)) {
                console.log(`Verification error: ${err.response?.data}`); // Debug log for error response
                toast.error(err.response?.data || 'Invalid OTP or verification failed');
            } else {
                console.log('Verification error: An error occurred'); // Debug log for general error
                toast.error('Invalid OTP or verification failed');
            }
        }
    };

    const handleGitHubLogin = () => {
        window.location.href = 'http://localhost:8081/auth/github';
    };

    return (
        <>
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
                        <div>
                            <button
                                type="submit"
                                className="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
                            >
                                {mode === 'login' ? 'Log In' : 'Sign Up'}
                            </button>
                        </div>
                        <div>
                            <GitHubLoginButton mode={mode} />
                        </div>
                    </>
                ) : (
                    <>
                        <div className="flex justify-center space-x-2">
                            {otp.map((digit, index) => (
                                <input
                                    key={index}
                                    type="text"
                                    value={digit}
                                    onChange={handleOtpChange(index)}
                                    maxLength={1}
                                    ref={el => {
                                        otpRefs.current[index] = el;
                                    }}
                                    className="w-10 px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm text-black text-center"
                                />
                            ))}
                        </div>
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
            <ToastContainer />
        </>
    );
};

export default AuthForm;
