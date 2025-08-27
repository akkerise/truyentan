import { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import { get } from '../../services/api';

interface Chapter {
  id: string;
  title: string;
  number: number;
}

interface ChapterListResponse {
  chapters: Chapter[];
  total: number;
}

export default function ChapterList({ novelId }: { novelId: string }) {
  const [chapters, setChapters] = useState<Chapter[]>([]);
  const [page, setPage] = useState(1);
  const [pageSize] = useState(10);
  const [total, setTotal] = useState(0);

  useEffect(() => {
    async function fetchChapters() {
      try {
        const data = await get<ChapterListResponse>(`/novels/${novelId}/chapters`, {
          params: { page, limit: pageSize },
        });
        setChapters(data.chapters);
        setTotal(data.total);
      } catch {
        setChapters([]);
        setTotal(0);
      }
    }
    fetchChapters();
  }, [novelId, page, pageSize]);

  const progress = localStorage.getItem(`readingProgress-${novelId}`);

  return (
    <div>
      {progress && (
        <div className="mb-4">
          <Link to={`/novels/${novelId}/chapters/${progress}`} className="text-blue-500 underline">
            Đọc tiếp tục
          </Link>
        </div>
      )}
      <ul>
        {chapters.map((chapter) => (
          <li key={chapter.id} className="py-1">
            <Link
              to={`/novels/${novelId}/chapters/${chapter.id}`}
              className="text-blue-500 hover:underline"
            >
              {chapter.number}. {chapter.title}
            </Link>
          </li>
        ))}
      </ul>
      <div className="flex gap-2 mt-4">
        <button
          type="button"
          disabled={page === 1}
          onClick={() => setPage((p) => Math.max(1, p - 1))}
          className="px-2 py-1 border rounded disabled:opacity-50"
        >
          Prev
        </button>
        <button
          type="button"
          disabled={page * pageSize >= total}
          onClick={() => setPage((p) => p + 1)}
          className="px-2 py-1 border rounded disabled:opacity-50"
        >
          Next
        </button>
      </div>
    </div>
  );
}
