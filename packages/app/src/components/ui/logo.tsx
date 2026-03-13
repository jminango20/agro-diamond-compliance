interface LogoProps {
  className?: string;
}

export function Logo({ className = 'h-8 w-8' }: LogoProps) {
  return (
    <svg
      viewBox="0 0 64 64"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
      className={className}
    >
      <defs>
        <linearGradient id="diamond-grad-1" x1="0" y1="0" x2="1" y2="1">
          <stop offset="0%" stopColor="#818cf8" />
          <stop offset="50%" stopColor="#6366f1" />
          <stop offset="100%" stopColor="#4f46e5" />
        </linearGradient>
        <linearGradient id="diamond-grad-2" x1="1" y1="0" x2="0" y2="1">
          <stop offset="0%" stopColor="#a5b4fc" />
          <stop offset="100%" stopColor="#6366f1" />
        </linearGradient>
        <linearGradient id="diamond-grad-3" x1="0.5" y1="0" x2="0.5" y2="1">
          <stop offset="0%" stopColor="#c7d2fe" />
          <stop offset="100%" stopColor="#818cf8" />
        </linearGradient>
        <linearGradient id="shield-grad" x1="0.5" y1="0" x2="0.5" y2="1">
          <stop offset="0%" stopColor="#ffffff" stopOpacity="0.25" />
          <stop offset="100%" stopColor="#ffffff" stopOpacity="0.05" />
        </linearGradient>
      </defs>

      {/* Diamond body — top-left facet */}
      <polygon points="32,4 12,24 32,28" fill="url(#diamond-grad-3)" />
      {/* Diamond body — top-right facet */}
      <polygon points="32,4 52,24 32,28" fill="url(#diamond-grad-2)" />
      {/* Diamond body — bottom-left facet */}
      <polygon points="12,24 32,28 32,58" fill="url(#diamond-grad-1)" />
      {/* Diamond body — bottom-right facet */}
      <polygon points="52,24 32,28 32,58" fill="#4f46e5" />

      {/* Facet edge highlights */}
      <line x1="32" y1="4" x2="32" y2="58" stroke="#ffffff" strokeOpacity="0.15" strokeWidth="0.5" />
      <line x1="12" y1="24" x2="52" y2="24" stroke="#ffffff" strokeOpacity="0.12" strokeWidth="0.5" />
      <line x1="12" y1="24" x2="32" y2="28" stroke="#ffffff" strokeOpacity="0.1" strokeWidth="0.5" />
      <line x1="52" y1="24" x2="32" y2="28" stroke="#ffffff" strokeOpacity="0.1" strokeWidth="0.5" />

      {/* Shield overlay — subtle security hint */}
      <path
        d="M32 16 C32 16, 22 20, 22 28 C22 38, 27 46, 32 50 C37 46, 42 38, 42 28 C42 20, 32 16, 32 16Z"
        fill="url(#shield-grad)"
        stroke="#ffffff"
        strokeOpacity="0.12"
        strokeWidth="0.6"
      />

      {/* Checkmark inside shield */}
      <polyline
        points="27,32 31,36 37,26"
        fill="none"
        stroke="#ffffff"
        strokeOpacity="0.5"
        strokeWidth="1.5"
        strokeLinecap="round"
        strokeLinejoin="round"
      />

      {/* Outer diamond outline glow */}
      <polygon
        points="32,4 12,24 32,58 52,24"
        fill="none"
        stroke="#818cf8"
        strokeOpacity="0.3"
        strokeWidth="1"
      />
    </svg>
  );
}

export function LogoFull({ className = 'h-8' }: LogoProps) {
  return (
    <div className={`flex items-center gap-2.5 ${className}`}>
      <Logo className="h-full w-auto" />
      <div className="flex flex-col leading-none">
        <span className="text-[15px] font-bold tracking-tight text-white">
          Diamond
        </span>
        <span className="text-[9px] font-medium tracking-[0.2em] text-indigo-400/80">
          ERC-3643
        </span>
      </div>
    </div>
  );
}
