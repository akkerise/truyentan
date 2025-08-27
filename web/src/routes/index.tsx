import { BrowserRouter, Routes, Route } from 'react-router-dom';
import Home from '../pages/Home';
import ChapterReader from '../pages/Reader/ChapterReader';

export default function AppRouter() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/reader/:id" element={<ChapterReader />} />
      </Routes>
    </BrowserRouter>
  );
}
