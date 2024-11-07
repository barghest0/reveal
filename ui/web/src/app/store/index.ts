import { configureStore } from '@reduxjs/toolkit'
import { productReducer } from 'entities/card-product/model/slice'
import { cartReducer } from 'entities/cart/model/slice'


export const store = configureStore({
  reducer: {
    cart: cartReducer,
    products: productReducer
  },
})

export type RootState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch