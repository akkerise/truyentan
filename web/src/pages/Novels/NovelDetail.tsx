import { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { get } from '../../services/api';
import ChapterList from '../../components/novels/ChapterList';

interface Novel {
  id: string;
  title: string;
  author: string;
  description: string;
  genres: string[];
  rating: number;
}

export default function NovelDetail() {
  const { id } = useParams<{ id: string }>();
  const [novel, setNovel] = useState<Novel | null>(null);

  useEffect(() => {
    async function fetchNovel() {
      if (!id) return;
      try {
        const data = await get<Novel>(`/novels/${id}`);
        setNovel(data);
      } catch {
        setNovel(null);
      }
    }
    fetchNovel();
  }, [id]);

  if (!novel) {
    return <div>Loading...</div>;
  }

  return (
    <div className="p-4">
      <h1 className="text-2xl font-bold mb-2">{novel.title}</h1>
      <p className="mb-2">Tác giả: {novel.author}</p>
      <p className="mb-4">{novel.description}</p>
      <p className="mb-2">Thể loại: {novel.genres.join(', ')}</p>
      <p className="mb-4">Đánh giá: {novel.rating}</p>
      <ChapterList novelId={novel.id} />
    </div>
  );
}
