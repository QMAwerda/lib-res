import React from 'react'

import { Post } from '../components/Post'
import { useParams } from 'react-router-dom'
import axios from '../axios'
import parse from 'html-react-parser'
import { posts } from '../data/index'

export const FullPost = () => {
	const { isbn } = useParams()
	const [data, setData] = React.useState()
	const [isLoading, setLoading] = React.useState(true)
	const [plainText, setPlainText] = React.useState('')

	React.useEffect(() => {
		axios
			.get(`/posts/${isbn}`)
			.then(res => {
				setData(res.data, data)
				setPlainText(parse(res.data.text))
				setLoading(false)
			})
			.catch(err => {
				alert('Ошибка просмотра статьи')
			})
	}, [])
	const obj = data

	return (
		<>
			<div
				style={{
					display: 'flex',
					justifyContent: 'center',
					alignItems: 'center',
				}}
			>
				<div
					style={{
						width: '80%',
					}}
				>
					<Post

						title={obj.title}
						author={obj.author}
						publisher={obj.publisher}
						createdAt={obj.created_at}
						isbn={obj.isbn}
						year_published={obj.year_published}
						description={obj.description}
						isFullPost
					>
						{plainText}
					</Post>
				</div>
			</div>
			\
		</>
	)
}
