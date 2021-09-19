/** @format */

import type { AppProps } from 'next/app';
import Nav from '../components/nav';
import '../styles/main.scss';

export default function App({ Component, pageProps }: AppProps) {
	return (
		<div>
			<Nav />
			<Component {...pageProps} />;
		</div>
	);
}
