/** @format */

import type { NextPage } from 'next';
import Head from 'next/head';

import classes from '../styles/home.module.scss';

const Home: NextPage = (): JSX.Element => {
	return (
		<div>
			<Head>
				<title>Cashresp - Home</title>
				<link rel='icon' href='/favicon.ico' />
			</Head>

			<main className={classes.main}>
				<div className={classes.content}>
					<h1 className={classes.h1}>
						Make <span>Money</span> By Sharing Your Opinions
					</h1>
					<button className={classes.cta}>Start Earning Today</button>
				</div>
				<img src='/index.svg' alt='' className={classes.svg} />
			</main>
		</div>
	);
};

export default Home;
