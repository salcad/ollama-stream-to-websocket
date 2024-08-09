'use client';

import 'highlight.js/styles/default.css'; 
import 'materialize-css/dist/css/materialize.min.css';

export default function ClientComponent({ children }: { children: React.ReactNode }) {
  return <>{children}</>;
}
