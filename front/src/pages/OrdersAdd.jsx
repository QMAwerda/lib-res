import React from 'react'

import { orders } from '../data/index'
import styles from '.'

export const OrdersAdd = () => {
	const emptyForm = {
		userId: '',
		isbn: '',
	}


	const [formData, setFormData] = React.useState(emptyForm)
	const handleSubmit = async data => {
		try {
			// await axios.post('/orders', data)
			alert('Успешно')
		} catch (error) {
			alert('Произошла ошибка при отправке')
			console.error(error)
		} finally {
			setFormData(emptyForm)
		}
	}
	const handleChange = e => {
		setFormData({ ...formData, [e.target.name]: e.target.value })
	}
	return (
		<>
			<form
				onSubmit={handleSubmit}
				style={{ display: 'grid', padding: '15px' }}
			>
				<label>
					фио
					<input
						type='text'
						name='id'
						minLength='2'
						maxLength='30'
						required
						onChange={handleChange}
					/>
				</label>
				<label>
					isbn
					<input
						type='text'
						name='isbm'
						minLength='2'
						maxLength='30'
						required
						value={formData.isbm}
						onChange={handleChange}
					/>
				</label>
			</form>
			<button style={{ padding: '10px' }} onSubmit={handleSubmit}>
				Добавить
			</button>
		</>
	)
}
