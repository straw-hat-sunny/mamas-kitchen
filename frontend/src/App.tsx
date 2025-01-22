// Code for the main App component
import './App.css'
import {Routes, Route} from 'react-router-dom';
import RecipePage from './pages/recipe';
import ListPage from './pages/list';


function App() {

  return (
    <>
      <Routes>
        <Route path="/" element={<ListPage />} />
        <Route path="/recipes" element={<ListPage />} />
        <Route path="/recipes/:id" element={<RecipePage />} />
      
      </Routes>
    </>
  )
}

export default App
