import React, { useState } from 'react';

const SLIDES = [
  "https://images.unsplash.com/photo-1504674900247-0877df9cc836?auto=format&fit=crop&w=800&q=80",
  "https://images.unsplash.com/photo-1555396273-367ea4eb4db5?auto=format&fit=crop&w=800&q=80",
  "https://images.unsplash.com/photo-1414235077428-338989a2e8c0?auto=format&fit=crop&w=800&q=80",
  "https://images.unsplash.com/photo-1559339352-11d035aa65de?auto=format&fit=crop&w=800&q=80",
  "https://images.unsplash.com/photo-1514933651103-005eec06c04b?auto=format&fit=crop&w=800&q=80",
  "https://images.unsplash.com/photo-1544025162-d76694265947?auto=format&fit=crop&w=800&q=80"
];

export default function ImageSlider() {
  const [currentSlide, setCurrentSlide] = useState(0);

  const nextSlide = () => setCurrentSlide((prev) => (prev + 1) % SLIDES.length);
  const prevSlide = () => setCurrentSlide((prev) => (prev - 1 + SLIDES.length) % SLIDES.length);

  return (
    <section id="gallery">
      <div className="container">
        <h2 className="section-title">Optic Buffers</h2>
        <div className="slider-container">
          <div
            className="slider"
            style={{ transform: `translateX(-${currentSlide * 100}%)`, transition: 'transform 0.4s ease' }}
          >
            {SLIDES.map((src, idx) => (
              <img key={idx} className="slide" src={src} alt={`Slide ${idx}`} />
            ))}
          </div>
          <button className="slider-btn prev" onClick={prevSlide}>◀</button>
          <button className="slider-btn next" onClick={nextSlide}>▶</button>
        </div>
      </div>
    </section>
  );
}