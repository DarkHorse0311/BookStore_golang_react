import React, { useEffect, useState } from 'react';
import Book from './BookContainer.js';
import '../style/Book.scss';

function BookPage() {
  let[loading, setLoading] = useState(true);
  let[book, setBook] = useState();
  let[gotSuggestions, setGotSuggestions] = useState(true);
  let[suggestions,setSuggestions] = useState([]);

  const getBook = async () => {
    const urlParams = new URLSearchParams(window.location.search);
    let idParam = urlParams.get('ID');
    let response = await fetch(`http://localhost:8080/books/`+idParam);
    let data = await response.json()
    
    setBook(data);
    setLoading(false);
   
  }

  const getSuggestions = async () => {

    let response = await fetch(`http://localhost:8080/books?genre=`+book.Book_genre);
    let data = await response.json()
    console.log(data);  
    setSuggestions(data);
    
    if(data.length >1){
        setGotSuggestions(false);
    }
  }
  useEffect(() => {
    if(loading){
        getBook();
    }
  });

  useEffect(() => {
        if(!loading){
            getSuggestions();
        }
    }, [loading])


  return (
    <div className="books-page-single">
      {
        loading ? "Loading..."
        :        
        <>
        <Book book={book}/>

        <div className="book-details">
            <div className="book-details-container">
            <div className="book-details-grid">
                <div className="book-details-heading"><h3>Details & Spec</h3></div>
                <hr></hr>
                <br></br>
                <div className="book-details-grid-row">
                    <div className="book-details-grid-row-cell">Title</div>
                    <div className="book-details-grid-row-cell">{book.Book_title}</div>
                </div>
                <div className="book-details-grid-row">
                    <div className="book-details-grid-row-cell">Genre</div>
                    <div className="book-details-grid-row-cell">{book.Book_genre}</div>
                </div>
                <div className="book-details-grid-row">
                    <div className="book-details-grid-row-cell">ISBN</div>
                    <div className="book-details-grid-row-cell">{book.ISBN}</div>
                </div>
                <div className="book-details-grid-row">
                    <div className="book-details-grid-row-cell">Publisher</div>
                    <div className="book-details-grid-row-cell">...</div>
                </div>
                <div className="book-details-grid-row">
                    <div className="book-details-grid-row-cell">Authors</div>
                    <div className="book-details-grid-row-cell">...</div>
                </div>
                <br></br>
                <br></br>  
                {
                     gotSuggestions ? ""
                     :
                        <>
                            <div className="book-details-heading"><h3>Suggestions</h3></div>
                            <hr></hr>
                            <br></br>
                        </>
                    
                }
            </div>
            </div>
        </div>
        </>
        }
        
        <div className="books-suggestions">
        <>
        {  
            gotSuggestions ? ""
            :       
                suggestions.filter(suggestion => suggestion.ISBN != book.ISBN).map( item => (
                    <Book book={item}/>
                ))     
        }
        </>
        </div>
    </div>
  );
}

export default BookPage;
