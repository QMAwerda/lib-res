import React from 'react'
import clsx from 'clsx'
import { Link } from 'react-router-dom'

import styles from './Post.module.scss'

import { Navigate } from 'react-router-dom'

import axios from '../../axios'

export const Post = ({
	title,
	author,
	publisher,
	isbn,
	year_published,
	description,
	isFullPost,
	isOrder,
}) => {
	// const [showConfirmation, setShowConfirmation] = React.useState(false)
	// const [confirmationResult, setConfirmationResult] = React.useState(null)
	// const handleClick = async () => {
	// 	try {
	// 		await axios.post('/posts', { post_id: id, user_id: isAuthorized })
	// 		alert('Успешно забронировано ')
	// 	} catch (err) {
	// 		if (
	// 			err.response.data ==
	// 			'duplicate key value violates unique constraint "posts_users_pkey"'
	// 		) {
	// 			setShowConfirmation(true)
	// 		} else {
	// 			console.error('Ошибка:', err) // Логируем другие ошибки
	// 		}
	// 	}
	// }
	// const handleConfirmation = res => {
	// 	setShowConfirmation(false)
	// 	setConfirmationResult(res) // Сохраняем результат
	// 	if (res) {
	// 		handleDelete() // Вызываем функцию удаления, если подтверждено
	// 	}
	// }
	// const handleDelete = async () => {
	// 	try {
	// 		console.log(id)
	// 		await axios.delete('/posts', {
	// 			params: { post_id: id, user_id: isAuthorized },
	// 		})
	// 		console.log({ user_id: isAuthorized, post_id: id })
	// 		alert('Бронь отменена')
	// 		document.location.reload()
	// 	} catch (error) {
	// 		console.error('Ошибка при отмене брони:', error)
	// 	}
	// }

	console.log(isbn)

	return (
		<div
			className={clsx(styles.root)}
			onClick={() => {
				return <Navigate to={`/posts/${isbn}`} />
			}}
		>
			<div className={styles.wrapper}>
				<div className={styles.indention}>
					{isOrder ? (
						<h2 className={clsx(styles.title)}>{title}</h2>
					) : (
						<h2 className={clsx(styles.title)}>
							<Link to={`/posts/${isbn}`}>{title}</Link>
						</h2>
					)}

					{isOrder ? (
						<div className={styles.description}>user_id {description}</div>
					) : (
						<div className={styles.description}>{description}</div>
					)}

					<ul className={styles.postDetails}>
						<li>{author}</li>
						<li>{publisher}</li>
						<li>{year_published}</li>
						<li>{isbn}</li>
					</ul>
				</div>
			</div>
			{isFullPost ? (
				<>
					<button className={styles.button}>Удалить</button>{' '}
					<button className={styles.button}>Изменить</button>
				</>
			) : (
				''
			)}
		</div>
	)
}
