import React, { useEffect, useState } from 'react';
import { BrowserRouter as Router, Route, Redirect, Switch, Link } from "react-router-dom";
import BooksPage from './components/BooksPage'
import BookPage from './components/BookPage'
import LoginPage from './components/LoginPage'
import CartPage from './components/CartPage'
import Header from './components/Header';
import './style/App.scss';


function App() {
  //login
  //signup
  //cart
  //orders
  //order
  //book
  //books

  //report pages tbd

  let [user, setUser] = useState();

  return (
    <div className="App">
      <Router>
          <Header user={user}/>
          <br></br>
          <br></br>
          <br></br>
          <br></br>
          <br></br>
          <Route path="/books" component={() => <BooksPage user={user} />} />
          <Route path="/book/:isbn" component={() => <BookPage user={user} />} />
          <Route path="/cart" component={() => <CartPage user={user}/>} />
          <Route path="/login" component={() => <LoginPage user={user} setUser={setUser}/>} />
          
       </Router>  
    </div>
  );
}

export default App;
