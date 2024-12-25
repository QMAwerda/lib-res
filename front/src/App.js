import React from 'react'
import Container from '@mui/material/Container'
import { Route, Routes } from 'react-router-dom'
import { Header } from './components'
import { Home, FullPost, Orders, OrdersAdd, BooksAdd } from './pages'

function App() {
	return (
		<>
			<Header />
			<Container maxWidth='lg'>
				<Routes>
					<Route path='/' element={<Home />} />
					<Route path='/posts/:isbn' element={<FullPost />} />
					<Route path='/orders' element={<Orders />} />
					<Route path='/orders-add' element={<OrdersAdd />} />
					<Route path='/books-add' element={<BooksAdd />} />
				</Routes>
			</Container>
		</>
	)
}

export default App
