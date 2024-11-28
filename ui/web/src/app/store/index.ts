import { configureStore } from '@reduxjs/toolkit'
import { productsCartReducer } from 'entities/product/model/slice'
import { userReducer } from 'entities/user'

export const store = configureStore({
  reducer: {
    productsCart: productsCartReducer,
    user: userReducer
  },
})

export type RootState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch