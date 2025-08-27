import { BrowserRouter, Routes, Route } from 'react-router-dom';
import Home from '../pages/Home';
import NovelList from '../pages/Novels/NovelList';

export default function AppRouter() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/novels" element={<NovelList />} />
      </Routes>
    </BrowserRouter>
  );
}
