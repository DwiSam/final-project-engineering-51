import React, { createContext, useState } from "react";

const AuthContext = createContext({
	token: "",
	isLoggedIn: false,
	login: (token) => {},
	logout: () => {},
});

export const AuthContextProvider = (props) => {
	const initialToken = localStorage.getItem("token");
	const [token, setToken] = useState(initialToken);
	const [isLoggedIn, setIsLoggedIn] = useState(false);

	const loginHandler = (token) => {
		console.log(token);
		setToken(token);
		localStorage.setItem("token", token);
		setIsLoggedIn(true);
	};

	const logoutHandler = () => {
		setToken(null);
		localStorage.removeItem("token");
	};

	const contextValue = {
		token: token,
		isLoggedIn: isLoggedIn,
		login: loginHandler,
		logout: logoutHandler,
	};

	return (
		<AuthContext.Provider value={contextValue}>
			{props.children}
		</AuthContext.Provider>
	);
};

export default AuthContext;