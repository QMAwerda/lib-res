import React from 'react'

import Grid from '@mui/material/Grid'
import { Link } from 'react-router-dom'
import { Post } from '../components/Post'
import { orders } from '../data/index'

export const Orders = () => {
	const [orderPosts, setOrderPosts] = React.useState(orders)
	const [isPostsLoading, setIsPostsLoading] = React.useState(false)
	const handleSubmit = () => {}
	return (
		<>
			<Link to='/orders-add'>
				<button style={{ padding: '10px' }}>Добавить</button>
			</Link>

			<Grid container spacing={4}>
				{(isPostsLoading ? [...Array(5)] : orderPosts).map((obj, index) =>
					isPostsLoading ? (
						<Grid key={index} item xs={12} sm={6} md={4} lg={4}>
							<Post key={index} isLoading={true} />
						</Grid>
					) : (
						<Grid key={index} item xs={12} sm={6} md={4} lg={4}>
							<Post
								id={obj.id}
								title={obj.isbn}
								year_published={obj.order_date}
								description={obj.userId}
								isOrder
							/>
						</Grid>
					)
				)}
			</Grid>
		</>
	)
}
