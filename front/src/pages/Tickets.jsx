import React from 'react'

import Grid from '@mui/material/Grid'

import { Post } from '../components/Post'

import axios from '../axios'

export const Tickets = () => {
	const userData = useSelector(state => state.auth.data)
	const [tickets, setTickets] = React.useState([])
	const [isPostsLoading, setIsPostsLoading] = React.useState(false)

	const fetchTickets = async () => {
		const result = await axios.get('/my-tickets', userData?.user_id)
		setTickets(result.data)
	}
	React.useEffect(() => {
		try {
			setIsPostsLoading(true)
			fetchTickets()
		} catch (err) {
			console.log(err)
		} finally {
			setIsPostsLoading(false)
		}
	}, [])
	if (!userData) {
		return
	}
	return (
		<>
			<Grid container spacing={4}>
				{(isPostsLoading ? [...Array(5)] : tickets).map((obj, index) =>
					isPostsLoading ? (
						<Grid key={index} item xs={12} sm={6} md={4} lg={4}>
							<Post key={index} isLoading={true} />
						</Grid>
					) : (
						<Grid key={index} item xs={12} sm={6} md={4} lg={4}>
							<Post
								id={obj.post_id}
								title={obj.title}
								imageUrl={obj.image_url}
								createdAt={obj.created_at}
								viewsCount={obj.views_count}
								tags={obj.tags ? obj.tags : ['..']}
								isAuthorized={userData?.user_id}
							/>
						</Grid>
					)
				)}
			</Grid>
		</>
	)
}
