import React from 'react'

import { Link } from 'react-router-dom'
import styles from './Header.module.scss'
import Container from '@mui/material/Container'

export const Header = () => {
	const defaultOptions = {
		loop: true,
		autoplay: true,
		rendererSettings: {
			preserveAspectRatio: 'xMidYMid slice',
		},
	}

	return (
		<div className={styles.root}>
			<Container maxWidth='lg'>
				<div className={styles.inner}>
					<Link className={styles.logo} to='/'>
						<p style={{ marginBottom: 0, color: 'red' }}>Библиотека</p>
					</Link>
					<div className={styles.buttons}>
						<Link to='/'>
							<button>главная</button>
						</Link>
						<Link to='/orders'>
							<button>заказы</button>
						</Link>
					</div>
				</div>
			</Container>
		</div>
	)
}
