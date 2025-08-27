import { BrowserRouter, Routes, Route } from 'react-router-dom';
import Home from '../pages/Home';
import NovelList from '../pages/Novels/NovelList';
import SignIn from '../pages/Auth/SignIn';
import SignUp from '../pages/Auth/SignUp';
import ChapterReader from '../pages/Reader/ChapterReader';

export default function AppRouter() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/novels" element={<NovelList />} />
        <Route path="/signin" element={<SignIn />} />
        <Route path="/signup" element={<SignUp />} />
        <Route path="/reader/:id" element={<ChapterReader />} />
      </Routes>
    </BrowserRouter>
  );
}
