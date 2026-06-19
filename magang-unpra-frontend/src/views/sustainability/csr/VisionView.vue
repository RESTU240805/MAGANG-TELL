<template>
  <!-- ═══ HERO ═══ -->
  <section class="vision-hero">
    <div class="vision-hero__bg-image"></div>
    <div class="vision-hero__bg-overlay"></div>
    <div class="vision-hero__bg-pattern"></div>
    <div class="vision-hero__content max-w-7xl mx-auto px-10">
      <div class="vision-hero__breadcrumb">
        <span>Sustainability</span>
        <span class="vision-hero__breadcrumb-sep">›</span>
        <span>CSR</span>
        <span class="vision-hero__breadcrumb-sep">›</span>
        <span class="vision-hero__breadcrumb-active">Vision & Mission</span>
      </div>
      <h1 class="vision-hero__title">Vision <span class="vision-hero__title-accent">And</span> Mission</h1>
      <p class="vision-hero__subtitle">
        Our commitment to sustainable development, ethical business practices,
        and community empowerment.
      </p>
      <div class="vision-hero__divider"></div>
    </div>
  </section>

  <!-- ═══ INTRO CARD ═══ -->
  <section class="vision-intro">
    <div class="max-w-7xl mx-auto px-10">
      <div class="vision-intro__card">
        <div class="vision-intro__card-accent"></div>
        <div class="vision-intro__card-body">
          <div class="vision-intro__icon-wrap">
            <svg width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round">
              <path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z"/>
            </svg>
          </div>
          <div class="vision-intro__card-text">
            <h2 class="vision-intro__card-title">Corporate Social Responsibility</h2>
            <p>
              Corporate Social Responsibility is defined as sustainable commitment to business to act ethically,
              operate legally and contribute to improving the economy while at the same time improving the quality
              of life of employees and their families, local communities and society as a whole around the company.
            </p>
            <p>
              The CSR program is an investment for companies for the growth and sustainability of the company and
              is no longer considered as a cost center, but as a profit center. The CSR program is the company's
              commitment to support the creation of sustainable development.
            </p>
            <p>
              The implementation of CSR programs is one form of implementation of the concept of Good Corporate
              Governance as a social and environmental responsibility as stipulated in the regulations / Law no.
              40 of 2007 concerning Limited Liability Companies.
            </p>
          </div>
        </div>
      </div>
    </div>
  </section>

  <!-- ═══ STRATEGY SECTION ═══ -->
  <section class="vision-strategy">
    <div class="max-w-7xl mx-auto px-10">
      <div class="vision-section-header">
        <div class="vision-section-header__tag">Strategy</div>
        <h2 class="vision-section-header__title">Strategy And Management Plan</h2>
        <p class="vision-section-header__desc">
          Key pillars guiding PT. Tanjungenim Lestari Pulp and Paper's CSR and governance approach.
        </p>
      </div>

      <div class="vision-strategy__grid">
        <div
          v-for="(item, index) in strategyItems"
          :key="index"
          class="vision-strategy__card"
        >
          <div class="vision-strategy__card-number">{{ String(index + 1).padStart(2, '0') }}</div>
          <div class="vision-strategy__card-bar"></div>
          <p class="vision-strategy__card-text">{{ item }}</p>
        </div>
      </div>
    </div>
  </section>

  <!-- ═══ OBJECTIVES SECTION ═══ -->
  <section class="vision-objectives">
    <div class="max-w-7xl mx-auto px-10">
      <div class="vision-objectives__layout">
        <!-- Left sticky label -->
        <div class="vision-objectives__sidebar">
          <div class="vision-objectives__sidebar-inner">
            <div class="vision-objectives__tag">Objectives</div>
            <h2 class="vision-objectives__title">
              The Objectives<br>of the CSR<br>Program
            </h2>
            <div class="vision-objectives__line"></div>
            <p class="vision-objectives__sidebar-note">
              PT. Tanjungenim Lestari<br>Pulp and Paper
            </p>
          </div>
        </div>

        <!-- Right content -->
        <div class="vision-objectives__content">
          <div class="vision-objectives__block">
            <div class="vision-objectives__block-icon">
              <svg width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round">
                <circle cx="12" cy="12" r="10"/><path d="M12 8v4l3 3"/>
              </svg>
            </div>
            <p>
              The purpose of implementing the Corporate Social Responsibility (CSR) of PT. Tanjungenim Lestari Pulp
              and Paper is not limited to fulfilling the Company's responsibilities as well as the form of
              compliance with regulations only. Furthermore, a well-implemented and systematic CSR program can form
              a harmonious, balanced relationship and support between the Company and the community in the surrounding.
            </p>
          </div>

          <div class="vision-objectives__block">
            <div class="vision-objectives__block-icon">
              <svg width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round">
                <polyline points="22 12 18 12 15 21 9 3 6 12 2 12"/>
              </svg>
            </div>
            <p>
              The company also believes that CSR programs are an important part that is inseparable from the
              benchmark of business success, apart from the implementation of good management and operational
              performance.
            </p>
          </div>

          <div class="vision-objectives__block vision-objectives__block--highlight">
            <div class="vision-objectives__block-icon">
              <svg width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round">
                <path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"/><circle cx="9" cy="7" r="4"/>
                <path d="M23 21v-2a4 4 0 0 0-3-3.87"/><path d="M16 3.13a4 4 0 0 1 0 7.75"/>
              </svg>
            </div>
            <p>
              On this basis, through the implementation of the responsibilities contained in the vision, mission
              and corporate values, the Company outlines CSR policies and programs that meet sustainability
              standards, have integrity, uphold business ethics, and comply with laws and regulations that apply.
              These standards are expected to create CSR programs that can have a positive, effective and targeted
              impact, be able to empower the community's capabilities, and be sustainable in the long term.
            </p>
          </div>
        </div>
      </div>
    </div>
  </section>


   <section class="contact-section">
      <div class="contact-container">
        <div class="contact-header anim-item">
          <span class="contact-label">GET IN TOUCH</span>
          <h2 class="contact-title">Our Offices</h2>
        </div>

        <div class="contact-grid">
          <div class="contact-card anim-item">
            <div class="contact-image">
              <img src="/images/jakarta.jpeg" alt="Jakarta Office" />
            </div>
            <div class="contact-body">
              <h3 class="contact-city">Jakarta</h3>
              <p class="contact-address">Menara Astra 22nd floor – Zona D, Jalan Jenderal Sudirman Kav. 5-6 Kel. Karet Tengsin, Kec. Tanah Abang</p>
              <a class="contact-phone" href="tel:+622186656809">
                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M22 16.92v3a2 2 0 0 1-2.18 2 19.79 19.79 0 0 1-8.63-3.07 19.5 19.5 0 0 1-6-6 19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72c.127.96.361 1.903.7 2.81a2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45c.907.339 1.85.573 2.81.7A2 2 0 0 1 22 16.92z"/></svg>
                +62 21 8665 6809 / 8665 6810
              </a>
            </div>
          </div>

          <div class="contact-card anim-item">
            <div class="contact-image">
              <img src="/images/lokasi pabrik.jpeg" alt="Mill Site" />
            </div>
            <div class="contact-body">
              <h3 class="contact-city">Mill Site</h3>
              <p class="contact-address">Desa Banuayu, Kec. Empat Petulai Dangku, Kab. Muara Enim, Sumatera Selatan</p>
              <a class="contact-phone" href="tel:+62713324150">
                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M22 16.92v3a2 2 0 0 1-2.18 2 19.79 19.79 0 0 1-8.63-3.07 19.5 19.5 0 0 1-6-6 19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72c.127.96.361 1.903.7 2.81a2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45c.907.339 1.85.573 2.81.7A2 2 0 0 1 22 16.92z"/></svg>
                (62) (713) 324150 – 324160
              </a>
            </div>
          </div>

          <div class="contact-card anim-item">
            <div class="contact-image">
              <img src="/images/palembang.jpeg" alt="Palembang Office" />
            </div>
            <div class="contact-body">
              <h3 class="contact-city">Palembang</h3>
              <p class="contact-address">Ruko Blok I/29, Komplek PTC Mall. Jl. R. Soekamto Palembang 30114, Sumatera Selatan</p>
              <a class="contact-phone" href="tel:+62711382409">
                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M22 16.92v3a2 2 0 0 1-2.18 2 19.79 19.79 0 0 1-8.63-3.07 19.5 19.5 0 0 1-6-6 19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72c.127.96.361 1.903.7 2.81a2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45c.907.339 1.85.573 2.81.7A2 2 0 0 1 22 16.92z"/></svg>
                (62) (711) 382409
              </a>
            </div>
          </div>

          <div class="contact-card anim-item">
            <div class="contact-image">
              <img src="/images/tarahan.jpeg" alt="Tarahan Port" />
            </div>
            <div class="contact-body">
              <h3 class="contact-city">Tarahan</h3>
              <p class="contact-address">Jl. Soekarno Hatta Km. 14, Batu Serampok Kel. Srengsem Kec. Panjang, Bandar Lampung</p>
              <a class="contact-phone" href="tel:+62721342311">
                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M22 16.92v3a2 2 0 0 1-2.18 2 19.79 19.79 0 0 1-8.63-3.07 19.5 19.5 0 0 1-6-6 19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72c.127.96.361 1.903.7 2.81a2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45c.907.339 1.85.573 2.81.7A2 2 0 0 1 22 16.92z"/></svg>
                (62) (721) 34231, 31318
              </a>
            </div>
          </div>
        </div>
      </div>
    </section>

    <section class="footer-info-section">
      <div class="footer-info-container">
        <div class="footer-info-grid">

          <div class="footer-info-col anim-item">
            <p class="footer-about-text">
              PT. Tanjungenim Lestari Pulp and Paper (PT. TEL), is one of the most exciting pulp mills in Indonesia today and the only pulp mill in the world to produce high-quality, bleached-hardwood Kraft pulp with 100 percent plantation grown Acacia mangium and Eucalyptus Pellita trees.
            </p>
          </div>

          <div class="footer-info-col footer-logos-col anim-item">
            <img src="/images/logosatu.jpeg" alt="Yayasan Pendidikan Tanjungenim Lestari - SMP, SD, PGTK Lematang Lestari" class="footer-logo-combined" />
          </div>

          <div class="footer-info-col footer-contact-col anim-item">
            <div class="footer-contact-row">
              <svg class="footer-contact-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M21 10c0 7-9 13-9 13s-9-6-9-13a9 9 0 1 1 18 0z"/><circle cx="12" cy="10" r="3"/></svg>
              <span>Kab. Muara Enim, Sumatera Selatan.</span>
            </div>
            <div class="footer-contact-row">
              <svg class="footer-contact-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M22 16.92v3a2 2 0 0 1-2.18 2 19.79 19.79 0 0 1-8.63-3.07 19.5 19.5 0 0 1-6-6 19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72c.127.96.361 1.903.7 2.81a2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45c.907.339 1.85.573 2.81.7A2 2 0 0 1 22 16.92z"/></svg>
              <span>(+62) 713-324-150</span>
            </div>
            <div class="footer-contact-row">
              <svg class="footer-contact-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="10"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 6v6l4 2"/></svg>
              <span>Mon-Fri: 8:00 – 17:00</span>
            </div>
          </div>

        </div>
      </div>
    </section>

  <footer class="site-footer">
      <div class="footer-container">
        <div class="footer-left-content"></div>
        <div class="footer-copyright">
          <p>Copyright 2026 PT TELPP. All right reserved.</p>
        </div>
      </div>
    </footer>
</template>

<script setup>
const strategyItems = [
  'To implement sound industrial relations Corporate Social Responsibility (CSR) and Good Corporate Governance',
  'To ensure a harmony industrial relations with stakeholders both internal and external',
  'To strengthen the implementation of Good Corporate Governance (GCG) in accordance with prevailing regulations and laws',
  'To do continuous Improvement Program (CIP) in each department',
]
</script>

<style scoped>
/* ─── HERO ─────────────────────────────────────────────── */
.vision-hero {
  position: relative;
  background: #0d1f17;
  padding: 4rem 0;
  overflow: hidden;
}
.vision-hero__bg-image {
  position: absolute;
  inset: 0;
  background: url('/images/backgroundour.jpeg') center/cover no-repeat;
}
.vision-hero__bg-overlay {
  position: absolute;
  inset: 0;
  background: linear-gradient(to right, rgba(13,31,23,0.92), rgba(13,31,23,0.7));
}
.vision-hero__bg-pattern {
  position: absolute;
  inset: 0;
  background-image:
    radial-gradient(circle at 70% 20%, rgba(34,197,94,0.12) 0%, transparent 55%),
    radial-gradient(circle at 10% 80%, rgba(16,185,129,0.08) 0%, transparent 45%);
  pointer-events: none;
}
.vision-hero__content { position: relative; z-index: 1; }

.vision-hero__breadcrumb {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 12px;
  color: rgba(255,255,255,0.45);
  letter-spacing: 0.06em;
  text-transform: uppercase;
  margin-bottom: 2rem;
}
.vision-hero__breadcrumb-sep { opacity: 0.4; }
.vision-hero__breadcrumb-active { color: #4ade80; }

.vision-hero__title {
  font-size: clamp(2.8rem, 6vw, 5rem);
  font-weight: 900;
  color: #ffffff;
  line-height: 1.05;
  letter-spacing: -0.02em;
  margin: 0 0 1.25rem;
}
.vision-hero__title-accent { color: #22c55e; }

.vision-hero__subtitle {
  font-size: 1.1rem;
  color: rgba(255,255,255,0.55);
  max-width: 480px;
  line-height: 1.7;
  margin: 0 0 2.5rem;
}
.vision-hero__divider {
  width: 60px;
  height: 3px;
  background: #22c55e;
  border-radius: 2px;
}

/* ─── INTRO CARD ────────────────────────────────────────── */
.vision-intro {
  background: #f8faf8;
  padding: 4rem 0;
}
.vision-intro__card {
  background: #ffffff;
  border: 1px solid #e2e8e2;
  border-radius: 16px;
  overflow: hidden;
  box-shadow: 0 4px 24px rgba(0,0,0,0.05);
  display: flex;
}
.vision-intro__card-accent {
  width: 5px;
  flex-shrink: 0;
  background: linear-gradient(180deg, #16a34a 0%, #4ade80 100%);
}
.vision-intro__card-body {
  padding: 2.5rem 3rem;
  display: flex;
  gap: 2rem;
  align-items: flex-start;
}
.vision-intro__icon-wrap {
  flex-shrink: 0;
  width: 56px;
  height: 56px;
  background: #f0fdf4;
  border: 1px solid #bbf7d0;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #16a34a;
}
.vision-intro__card-text { flex: 1; }
.vision-intro__card-title {
  font-size: 0.7rem;
  font-weight: 700;
  color: #16a34a;
  text-transform: uppercase;
  letter-spacing: 0.12em;
  margin: 0 0 1.25rem;
}
.vision-intro__card-text p {
  font-size: 0.95rem;
  color: #4b5563;
  line-height: 1.8;
  margin: 0 0 0.9rem;
}
.vision-intro__card-text p:last-child { margin-bottom: 0; }

/* ─── STRATEGY ──────────────────────────────────────────── */
.vision-strategy {
  background: #0d1f17;
  padding: 5rem 0;
}
.vision-section-header {
  text-align: center;
  margin-bottom: 3.5rem;
}
.vision-section-header__tag {
  display: inline-block;
  background: rgba(34,197,94,0.12);
  color: #4ade80;
  font-size: 0.68rem;
  font-weight: 700;
  letter-spacing: 0.14em;
  text-transform: uppercase;
  padding: 4px 14px;
  border-radius: 20px;
  border: 1px solid rgba(74,222,128,0.25);
  margin-bottom: 1rem;
}
.vision-section-header__title {
  font-size: clamp(1.6rem, 3vw, 2.4rem);
  font-weight: 800;
  color: #ffffff;
  margin: 0 0 0.75rem;
  letter-spacing: -0.01em;
}
.vision-section-header__desc {
  font-size: 0.95rem;
  color: rgba(255,255,255,0.45);
  max-width: 480px;
  margin: 0 auto;
  line-height: 1.6;
}

.vision-strategy__grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(240px, 1fr));
  gap: 1.25rem;
}
.vision-strategy__card {
  background: rgba(255,255,255,0.04);
  border: 1px solid rgba(255,255,255,0.08);
  border-radius: 14px;
  padding: 2rem 1.75rem;
  transition: background 0.2s, border-color 0.2s, transform 0.2s;
  cursor: default;
}
.vision-strategy__card:hover {
  background: rgba(34,197,94,0.06);
  border-color: rgba(74,222,128,0.3);
  transform: translateY(-3px);
}
.vision-strategy__card-number {
  font-size: 2.5rem;
  font-weight: 900;
  color: rgba(74,222,128,0.2);
  line-height: 1;
  margin-bottom: 0.75rem;
  letter-spacing: -0.03em;
}
.vision-strategy__card-bar {
  width: 28px;
  height: 2px;
  background: #22c55e;
  border-radius: 2px;
  margin-bottom: 1rem;
}
.vision-strategy__card-text {
  font-size: 0.9rem;
  color: rgba(255,255,255,0.72);
  line-height: 1.7;
  margin: 0;
}

/* ─── OBJECTIVES ────────────────────────────────────────── */
.vision-objectives {
  background: #f8faf8;
  padding: 5rem 0;
}
.vision-objectives__layout {
  display: grid;
  grid-template-columns: 280px 1fr;
  gap: 4rem;
  align-items: start;
}
.vision-objectives__sidebar-inner {
  position: sticky;
  top: 6rem;
}
.vision-objectives__tag {
  display: inline-block;
  background: #f0fdf4;
  color: #16a34a;
  font-size: 0.65rem;
  font-weight: 700;
  letter-spacing: 0.14em;
  text-transform: uppercase;
  padding: 4px 12px;
  border-radius: 20px;
  border: 1px solid #bbf7d0;
  margin-bottom: 1rem;
}
.vision-objectives__title {
  font-size: 2rem;
  font-weight: 900;
  color: #111827;
  line-height: 1.2;
  letter-spacing: -0.02em;
  margin: 0 0 1.5rem;
}
.vision-objectives__line {
  width: 36px;
  height: 3px;
  background: #22c55e;
  border-radius: 2px;
  margin-bottom: 1.5rem;
}
.vision-objectives__sidebar-note {
  font-size: 0.8rem;
  color: #9ca3af;
  line-height: 1.6;
  margin: 0;
}

.vision-objectives__content {
  display: flex;
  flex-direction: column;
  gap: 1.25rem;
}
.vision-objectives__block {
  background: #ffffff;
  border: 1px solid #e5e7eb;
  border-radius: 14px;
  padding: 1.75rem 2rem;
  display: flex;
  gap: 1.25rem;
  align-items: flex-start;
  transition: box-shadow 0.2s, border-color 0.2s;
}
.vision-objectives__block:hover {
  border-color: #bbf7d0;
  box-shadow: 0 4px 20px rgba(34,197,94,0.07);
}
.vision-objectives__block--highlight {
  border-color: #bbf7d0;
  background: #f0fdf4;
}
.vision-objectives__block-icon {
  flex-shrink: 0;
  width: 42px;
  height: 42px;
  background: #f0fdf4;
  border: 1px solid #bbf7d0;
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #16a34a;
}
.vision-objectives__block--highlight .vision-objectives__block-icon {
  background: #dcfce7;
  border-color: #86efac;
}
.vision-objectives__block p {
  font-size: 0.92rem;
  color: #4b5563;
  line-height: 1.8;
  margin: 0;
  padding-top: 8px;
}

/* ─── FOOTER BAR ─────────────────────────────────────────── */
.vision-footer-bar {
  height: 5px;
  background: #e5e7eb;
}
.vision-footer-bar__fill {
  height: 100%;
  width: 100%;
  background: linear-gradient(90deg, #15803d 0%, #22c55e 50%, #86efac 100%);
}

/* ─── RESPONSIVE ─────────────────────────────────────────── */
@media (max-width: 768px) {
  .vision-hero { padding: 5rem 0 3.5rem; }
  .vision-intro__card-body { flex-direction: column; gap: 1.25rem; padding: 1.75rem; }
  .vision-objectives__layout { grid-template-columns: 1fr; gap: 2rem; }
  .vision-objectives__sidebar-inner { position: static; }
  .vision-objectives__title { font-size: 1.6rem; }
}


.contact-container {
  max-width: 1140px;
  margin: 0 auto;
  padding: 0 40px;
}

.contact-header {
  text-align: center;
  margin-bottom: 48px;
}

.contact-label {
  color: #5F9E42;
  font-size: 12px;
  font-weight: 700;
  letter-spacing: 0.25em;
  text-transform: uppercase;
  display: block;
  margin-bottom: 10px;
}

.contact-title {
  font-size: 32px;
  font-weight: 800;
  color: #18301f;
  letter-spacing: -0.02em;
}

.contact-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 24px;
}

@media (max-width: 1024px) {
  .contact-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (max-width: 640px) {
  .contact-grid {
    grid-template-columns: 1fr;
  }
}

.contact-card {
  background: #ffffff;
  border: 1px solid #dcecd2;
  border-radius: 20px;
  padding: 28px 24px;
  transition: all 0.4s cubic-bezier(0.16, 1, 0.3, 1);
}

.contact-card:hover {
  transform: translateY(-6px);
  border-color: #a9d68c;
  box-shadow: 0 16px 32px rgba(95, 158, 66, 0.12);
}

.contact-icon {
  width: 44px;
  height: 44px;
  border-radius: 12px;
  background: #eaf6e1;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 18px;
  color: #5F9E42;
}

.contact-icon svg {
  width: 22px;
  height: 22px;
}

.contact-city {
  font-size: 17px;
  font-weight: 700;
  color: #18301f;
  margin-bottom: 10px;
}

.contact-address {
  font-size: 13px;
  color: #6b7b66;
  line-height: 1.6;
  margin-bottom: 16px;
  font-weight: 400;
}

.contact-phone {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 13.5px;
  font-weight: 600;
  color: #3b6d11;
  text-decoration: none;
  transition: color 0.3s ease;
}

.contact-phone:hover {
  color: #27500a;
}

.contact-phone svg {
  width: 15px;
  height: 15px;
  flex-shrink: 0;
}

.contact-section {
  background: linear-gradient(180deg, #ffffff 0%, #eef7e9 100%);
  padding: 80px 0 70px;
}

.contact-container {
  max-width: 1140px;
  margin: 0 auto;
  padding: 0 40px;
}

.contact-header {
  text-align: center;
  margin-bottom: 48px;
}

.contact-label {
  color: #5F9E42;
  font-size: 12px;
  font-weight: 700;
  letter-spacing: 0.25em;
  text-transform: uppercase;
  display: block;
  margin-bottom: 10px;
}

.contact-title {
  font-size: 32px;
  font-weight: 800;
  color: #18301f;
  letter-spacing: -0.02em;
}

.contact-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 24px;
}

@media (max-width: 1024px) {
  .contact-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (max-width: 640px) {
  .contact-grid {
    grid-template-columns: 1fr;
  }
}

.contact-card {
  background: #ffffff;
  border: 1px solid #dcecd2;
  border-radius: 20px;
  overflow: hidden;
  transition: all 0.4s cubic-bezier(0.16, 1, 0.3, 1);
}

.contact-card:hover {
  transform: translateY(-6px);
  border-color: #a9d68c;
  box-shadow: 0 16px 32px rgba(95, 158, 66, 0.15);
}

.contact-image {
  width: 100%;
  aspect-ratio: 4 / 3;
  overflow: hidden;
  background: #eaf6e1;
}

.contact-image img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  display: block;
  transition: transform 0.6s cubic-bezier(0.16, 1, 0.3, 1);
}

.contact-card:hover .contact-image img {
  transform: scale(1.06);
}

.contact-body {
  padding: 20px 18px 22px;
}

.contact-city {
  font-size: 18px;
  font-weight: 800;
  color: #1a4d8f;
  margin-bottom: 12px;
  letter-spacing: -0.01em;
}

.contact-address {
  font-size: 13px;
  color: #5a6b56;
  line-height: 1.6;
  margin-bottom: 16px;
  font-weight: 400;
}

.contact-phone {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 13.5px;
  font-weight: 500;
  color: #2c3a2a;
  text-decoration: none;
  transition: color 0.3s ease;
}

.contact-phone:hover {
  color: #3b6d11;
}

.contact-phone svg {
  width: 15px;
  height: 15px;
  flex-shrink: 0;
  color: #5F9E42;
}
/* Responsif untuk Layar Handphone (Mobile) */
@media (max-width: 768px) {
  .footer-container {
    flex-direction: column;
    gap: 15px;
    padding: 0 20px;
  }
  
  .footer-copyright {
    text-align: center; /* Di HP otomatis pindah ke tengah bawah agar rapi */
    width: 100%;
  }
  
  .footer-copyright p {
    font-size: 12px; /* Ukuran sedikit disesuaikan untuk layar kecil */
  }
}

.footer-info-section {
  background: #ffffff;
  border-top: 1px solid #e4efdb;
  padding: 56px 0;
}

.footer-info-container {
  max-width: 1140px;
  margin: 0 auto;
  padding: 0 40px;
}
.footer-info-grid {
  display: grid;
  grid-template-columns: 1fr 1.1fr 0.9fr;
  gap: 48px;
  align-items: start;
}

.footer-info-col {
  display: flex;
  flex-direction: column;
}

.footer-about-text {
  margin: 0;
  font-size: 14px;
  color: #7a8a76;
  line-height: 1.75;
  font-weight: 400;
}

.footer-logos-col {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 0;
  margin: 0;
  padding: 0;
}

.footer-logos-col {
  display: flex;
  align-items: flex-start;
  justify-content: center;
  margin: 0;
  padding: 0;
}

.footer-logo-combined {
  width: 100%;
  max-width: 380px;
  height: auto;
  object-fit: contain;
  display: block;
}

.footer-contact-col {
  display: flex;
  flex-direction: column;
  gap: 20px;
  max-width: 280px;
  margin: 0;
  padding: 0;
}

.footer-contact-row {
  display: flex;
  align-items: center;
  gap: 14px;
  font-size: 14px;
  color: #5a6b56;
  font-weight: 400;
  line-height: 1.4;
  margin: 0;
}
.footer-contact-icon {
  width: 19px;
  height: 19px;
  flex-shrink: 0;
  color: #5F9E42;
}
/* ============ FOOTER ============ */
.site-footer {
  width: 100%;
  /* Warna hijau tua disesuaikan dengan identitas PT TEL pada gambar */
  background-color: #5F9E42; 
  padding: 25px 0; /* Memberikan ruang tinggi baris yang pas */
  margin-top: 40px; /* Jarak pemisah dari konten di atasnya */
}

.footer-container {
  display: flex;
  align-items: center;
  justify-content: center; /* DIUBAH: Membuat konten berkumpul di tengah */
  max-width: 1140px; /* Menyelaraskan lebar dengan konten atas website */
  margin: 0 auto;
  padding: 0 40px; /* Jarak aman agar teks tidak terlalu menempel ke pinggir layar */
}

/* Kolom Kiri Footer */
.footer-left-content {
  display: none; /* DIUBAH: Disembunyikan karena kosong, agar tidak memakan space */
}

/* Kolom Kanan Footer (Teks Copyright) */
.footer-copyright {
  text-align: center; /* DIUBAH: Memastikan teks rata tengah */
  width: 100%; /* Memastikan block mengambil ruang penuh untuk fungsi centering */
}

.footer-copyright p {
  font-family: Arial, sans-serif;
  font-size: 13px; /* Ukuran teks kecil, tipis, dan normal untuk footer */
  color: #ffffff;  /* Warna teks putih bersih agar kontras dengan background hijau */
  margin: 0;
  font-weight: normal;
  letter-spacing: 0.3px;
}
</style>