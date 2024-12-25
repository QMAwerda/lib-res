import React from 'react'
import axios from '../axios'


export const BooksAdd = () => {
	const emptyForm = {
		isbn: '',
		title: '',
		author: '',
		publisher: '',
		year_published: '',
		amount: '',
		description: '',
	}


	const [formData, setFormData] = React.useState(emptyForm)
	const handleSubmit = async () => {
		try {
			console.log(formData)
			await axios.post('/books', formData, {
				headers: {
					'Content-Type': 'application/json'
				}
			})
			alert('Успешно')
		} catch (error) {
			alert(error)
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
					Название книги
					<input
						type='text'
						name='title'
						minLength='2'
						maxLength='30'
						required
						onChange={handleChange}
					/>
				</label>
				<label>
					Автор
					<input
						type='text'
						name='author'
						minLength='2'
						maxLength='30'
						required
						onChange={handleChange}
					/>
				</label>
				<label>
					Издатель
					<input
						type='text'
						name='publisher'
						minLength='2'
						maxLength='30'
						required
						onChange={handleChange}
					/>
				</label>
				<label>
					Описание
					<input
						type='text'
						name='description'
						minLength='2'
						maxLength='30'
						required
						onChange={handleChange}
					/>
				</label>
				<label>
					Год выпуска
					<input
						type='text'
						name='year_published'
						minLength='2'
						maxLength='30'
						required
						onChange={handleChange}
					/>
				</label>
				<label>
					Количество
					<input
						type='text'
						name='amount'
						minLength='1'
						maxLength='3'
						required
						onChange={handleChange}
					/>
				</label>
				<label>
					isbn
					<input
						type='text'
						name='isbn'
						minLength='2'
						maxLength='30'
						required
						value={formData.isbm}
						onChange={handleChange}
					/>
				</label>
				<button style={{ padding: '10px' }}>
					Добавить
				</button>
			</form>

		</>
	)
}
