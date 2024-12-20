import React from "react";
import ReactDOM from 'react-dom/client';
import { App } from "./app/app";
import { Provider } from "react-redux";
import { store } from "app/store";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import { MainPage } from "pages/main";
import { CartPage } from "pages/cart";
import './index.css'

const container = document.getElementById('root');

const root = ReactDOM.createRoot(container as HTMLDivElement);

root.render(
    <React.StrictMode>
        <BrowserRouter>
            <Provider store={store}>
                <Routes>
                    <Route path="/" element={<App/>}/>
                    <Route path="mainPage" element={<MainPage/>}/>
                    <Route path="cartPage" element={<CartPage/>}/> 
                </Routes>
            </Provider>
        </BrowserRouter>
    </React.StrictMode>
);

