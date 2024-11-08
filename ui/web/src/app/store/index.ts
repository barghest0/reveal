import { configureStore } from '@reduxjs/toolkit'
import { productsCartReducer } from 'entities/cart/model/slice'

export const store = configureStore({
  reducer: {
    productsCart: productsCartReducer,
  },
})

export type RootState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch