import { useCallback, useEffect, useRef, useState } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import { get, post } from '../../services/api';

interface Chapter {
  id: number;
  novel_id: number;
  title: string;
  content: string;
  prev_id?: number;
  next_id?: number;
}

interface Progress {
  novelId: number;
  chapterId: number;
  positionOffset: number;
}

export default function ChapterReader() {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();
  const [chapter, setChapter] = useState<Chapter | null>(null);
  const saveTimer = useRef<number>();

  useEffect(() => {
    if (!id) return;
    get<Chapter>(`/chapters/${id}`).then(setChapter);
  }, [id]);

  const saveProgress = useCallback(
    (position: number) => {
      if (!chapter) return;
      const body = {
        novelId: chapter.novel_id,
        chapterId: chapter.id,
        positionOffset: position,
      };
      post('/me/progress', body).catch(() => {});
    },
    [chapter],
  );

  const handleScroll = useCallback(() => {
    const position = window.scrollY;
    window.clearTimeout(saveTimer.current);
    saveTimer.current = window.setTimeout(() => saveProgress(position), 500);
  }, [saveProgress]);

  useEffect(() => {
    window.addEventListener('scroll', handleScroll);
    return () => {
      window.removeEventListener('scroll', handleScroll);
      window.clearTimeout(saveTimer.current);
    };
  }, [handleScroll]);

  useEffect(() => {
    const handleKey = (e: KeyboardEvent) => {
      if ((e.key === 'ArrowRight' || e.key === 'PageDown') && chapter?.next_id) {
        navigate(`/reader/${chapter.next_id}`);
      } else if ((e.key === 'ArrowLeft' || e.key === 'PageUp') && chapter?.prev_id) {
        navigate(`/reader/${chapter.prev_id}`);
      }
    };
    window.addEventListener('keydown', handleKey);
    return () => window.removeEventListener('keydown', handleKey);
  }, [chapter, navigate]);

  useEffect(() => {
    if (!chapter) return;
    get<Progress>(`/me/progress/${chapter.novel_id}`)
      .then((p) => {
        if (p.chapterId === chapter.id) {
          window.scrollTo({ top: p.positionOffset });
        }
      })
      .catch(() => {});
  }, [chapter]);

  const goPrev = useCallback(() => {
    if (chapter?.prev_id) {
      navigate(`/reader/${chapter.prev_id}`);
    }
  }, [chapter, navigate]);

  const goNext = useCallback(() => {
    if (chapter?.next_id) {
      navigate(`/reader/${chapter.next_id}`);
    }
  }, [chapter, navigate]);

  if (!chapter) {
    return <div>Loading...</div>;
  }

  return (
    <div className="p-4">
      <h1 className="mb-4 text-xl font-bold">{chapter.title}</h1>
      <div className="whitespace-pre-line" dangerouslySetInnerHTML={{ __html: chapter.content }} />
      <div className="mt-8 flex justify-between">
        <button disabled={!chapter.prev_id} onClick={goPrev}>
          Prev
        </button>
        <button disabled={!chapter.next_id} onClick={goNext}>
          Next
        </button>
      </div>
    </div>
  );
}
