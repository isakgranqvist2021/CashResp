/** @format */

import { useState, useRef, useEffect } from 'react';
import classes from '../styles/modules/hero.module.scss';
import DOTS from 'vanta/dist/vanta.dots.min';

export default function Hero(): JSX.Element {
	const [vantaEffect, setVantaEffect] = useState<any>(0);
	const myRef = useRef(null);

	useEffect(() => {
		if (!vantaEffect) {
			setVantaEffect(
				DOTS({
					el: myRef.current,
					mouseControls: true,
					touchControls: true,
					gyroControls: false,
					minHeight: 200.0,
					minWidth: 200.0,
					scale: 1.0,
					scaleMobile: 1.0,
					color: 0x8fa181,
					color2: 0x7f9172,
					size: 1.3,
					spacing: 31.0,
				})
			);
		}
		return () => {
			if (vantaEffect) vantaEffect.destroy();
		};
	}, [vantaEffect]);

	return (
		<div ref={myRef} className={classes.hero}>
			<div className={classes.content}>
				<h1 className={classes.h1}>
					Make <span>Money</span> By Sharing Your Opinions
				</h1>
				<button className={classes.cta}>Start Earning Today</button>
			</div>
		</div>
	);
}
