// src/components/Header.tsx

import React from 'react';
import Image from 'next/image';
import logo from '../app/logo_bw_.svg';

const Header: React.FC = () => {
    return (
        <div className="flex items-center justify-center mb-8">
            <Image src={logo} alt="Logo" width={300} height={70} />
        </div>
    );
};

export default Header;
