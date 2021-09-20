/** @format */

import type { NextPage } from 'next';
import Head from 'next/head';
import Hero from '../components/hero';
import classes from '../styles/modules/home.module.scss';

const Home: NextPage = (): JSX.Element => {
	return (
		<div>
			<Head>
				<title>Cashresp - Home</title>
				<link rel='icon' href='/favicon.ico' />
			</Head>

			<main className={classes.main}>
				<Hero />
			</main>

			<script src='/three.min.js' />
		</div>
	);
};

export default Home;
