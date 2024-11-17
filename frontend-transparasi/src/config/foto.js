export const BASE_URL = os.env.BE_URL || 'http://localhost:8000';
export const getImageUrl = (path) => {
  if (!path) return '';
  if (path.startsWith('http')) return path;
  const cleanPath = path.startsWith('/') ? path.slice(1) : path;
  const url = `${BASE_URL}/${cleanPath}`;
  return url;
};
