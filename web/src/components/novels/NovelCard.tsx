interface Genre {
  id?: number;
  name: string;
}

export interface Novel {
  id: number;
  title: string;
  author: string;
  description: string;
  status: string;
  genres?: Genre[];
}

interface NovelCardProps {
  novel: Novel;
}

export default function NovelCard({ novel }: NovelCardProps) {
  return (
    <div className="border rounded p-4 shadow-sm space-y-1">
      <h2 className="text-lg font-semibold">{novel.title}</h2>
      <p className="text-sm text-gray-600">{novel.author}</p>
      <p className="text-sm text-gray-700">
        {novel.description.length > 100
          ? novel.description.slice(0, 100) + '...'
          : novel.description}
      </p>
      {novel.genres && (
        <div className="flex flex-wrap gap-1 mt-2">
          {novel.genres.map((g) => (
            <span key={g.id ?? g.name} className="bg-gray-200 text-xs px-2 py-0.5 rounded">
              {g.name}
            </span>
          ))}
        </div>
      )}
    </div>
  );
}

export function NovelCardSkeleton() {
  return (
    <div className="border rounded p-4 shadow-sm animate-pulse space-y-2">
      <div className="h-6 bg-gray-200 rounded w-3/4" />
      <div className="h-4 bg-gray-200 rounded w-1/2" />
      <div className="h-4 bg-gray-200 rounded w-full" />
      <div className="h-4 bg-gray-200 rounded w-5/6" />
    </div>
  );
}
