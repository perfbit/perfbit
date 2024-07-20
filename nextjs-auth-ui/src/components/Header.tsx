// src/components/Header.tsx

import React from 'react';
import Image from 'next/image';
import logo from '../app/logo.svg';
import siteIcon from '../app/siteicon.svg';

const Header: React.FC = () => {
    return (
        <div className="flex items-center justify-center mb-8">
            <Image src={siteIcon} alt="Site Icon" width={50} height={50} />
            <Image src={logo} alt="Logo" width={200} height={50} className="ml-4" />
        </div>
    );
};

export default Header;
