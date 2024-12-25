import React from 'react'
import axios from '../axios'
import { useParams } from 'react-router-dom'
import { Post } from '../components/Post/index'
import Grid from '@mui/material/Grid'
import { useSelector } from 'react-redux'

export const TagPosts = () => {
	const userData = useSelector(state => state.auth.data)
	const { name } = useParams()
	const [posts, setPosts] = React.useState()
	const [isLoading, setLoading] = React.useState(true)

	React.useEffect(() => {
		console.log('starts')
		axios
			.get(`/tags/${name}`)
			.then(res => {
				setPosts(res.data)
				setLoading(false)
			})
			.catch(err => {
				console.warn(err)
			})
	}, [])

	return (
		<>
			<Grid xs={8} item>
				{(isLoading ? [...Array(5)] : posts).map((obj, index) =>
					isLoading ? (
						<Post key={index} isLoading={true} />
					) : (
						<Post
							id={obj.post_id}
							title={obj.title}
							imageUrl={obj.image_url}
							createdAt={obj.created_at}
							viewsCount={obj.views_count}
							tags={obj.tags}
							isAuthorized={userData?.user_id}
						/>
					)
				)}
			</Grid>
		</>
	)
}
