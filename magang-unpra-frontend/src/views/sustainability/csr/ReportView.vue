<template>
  <!-- ═══ HERO ═══ -->
  <section class="csrr-hero">
    <div class="csrr-hero__bg"></div>
    <div class="csrr-hero__content max-w-7xl mx-auto px-10">
      <div class="csrr-hero__breadcrumb">
        <span>Sustainability</span>
        <span class="csrr-sep">›</span>
        <span>CSR</span>
        <span class="csrr-sep">›</span>
        <span class="csrr-active">CSR Report</span>
      </div>
      <h1 class="csrr-hero__title">CSR <span class="csrr-hero__accent">Report</span></h1>
      <p class="csrr-hero__sub">Annual quarterly sustainability reports by PT. Tanjungenim Lestari Pulp and Paper.</p>
      <div class="csrr-hero__divider"></div>
    </div>
  </section>

  <!-- ═══ REPORT LIST ═══ -->
  <section class="csrr-body">
    <div class="max-w-7xl mx-auto px-10">

      <div v-for="year in reports" :key="year.year" class="csrr-year-block">

        <!-- Year Header -->
        <div class="csrr-year-header">
          <div class="csrr-year-badge">{{ year.year }}</div>
          <div class="csrr-year-line"></div>
          <span class="csrr-year-label">CSR Report Fiscal Year {{ year.year }}</span>
        </div>

        <!-- Quarter Cards -->
        <div class="csrr-quarters">
          <div
            v-for="q in year.quarters"
            :key="q.label"
            class="csrr-card"
            :class="q.file ? 'csrr-card--available' : 'csrr-card--unavailable'"
            @click="q.file && openPdf(q.file)"
          >
            <!-- Icon -->
            <div class="csrr-card__icon">
              <svg width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round">
                <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/>
                <polyline points="14 2 14 8 20 8"/>
                <line v-if="q.file" x1="12" y1="18" x2="12" y2="12"/>
                <line v-if="q.file" x1="9" y1="15" x2="15" y2="15"/>
                <line v-else x1="9" y1="15" x2="15" y2="15"/>
              </svg>
            </div>

            <!-- Info -->
            <div class="csrr-card__info">
              <p class="csrr-card__quarter">{{ q.label }}</p>
              <p class="csrr-card__period">{{ q.period }}</p>
            </div>

            <!-- Action -->
            <div class="csrr-card__action">
              <span v-if="q.file" class="csrr-card__btn">
                <svg width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
                  <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"/>
                  <polyline points="7 10 12 15 17 10"/>
                  <line x1="12" y1="15" x2="12" y2="3"/>
                </svg>
                View PDF
              </span>
              <span v-else class="csrr-card__na">Not Available</span>
            </div>
          </div>
        </div>

      </div>
    </div>
  </section>

  <!-- ═══ PDF MODAL ═══ -->
  <Teleport to="body">
    <div v-if="activePdf" class="csrr-modal" @click.self="activePdf = null">
      <div class="csrr-modal__box">
        <div class="csrr-modal__header">
          <div class="csrr-modal__header-left">
            <div class="csrr-modal__pdf-icon">
              <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/>
                <polyline points="14 2 14 8 20 8"/>
              </svg>
            </div>
            <span class="csrr-modal__title">{{ activePdfName }}</span>
          </div>
          <button class="csrr-modal__close" @click="activePdf = null">
            <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><line x1="18" y1="6" x2="6" y2="18"/><line x1="6" y1="6" x2="18" y2="18"/></svg>
          </button>
        </div>
        <iframe :src="activePdf" class="csrr-modal__iframe"></iframe>
      </div>
    </div>
  </Teleport>

  <div class="csrr-footer-bar"></div>
</template>

<script setup>
import { ref } from 'vue'

const activePdf = ref(null)
const activePdfName = ref('')

const reports = [
  {
    year: 2018,
    quarters: [
      { label: '1st Quarter', file: null },
      { label: '2nd Quarter', file: null },
      { label: '3rd Quarter', file: null },
      { label: '4th Quarter', file: null },
    ]
  },
  {
    year: 2019,
    quarters: [
      { label: '1st Quarter', period: 'Jan – Mar 2019', file: 'CSRReport2019_1st_Quarter_Apr_Ju.pdf' },
      { label: '2nd Quarter', file: null },
      { label: '3rd Quarter', file: null },
      { label: '4th Quarter', period: 'jan – mar 2019', file: 'CSR-Report-2019-4th-Quarter.pdf' },
    ]
  },
  {
    year: 2020,
    quarters: [
      { label: '1st Quarter', period: 'Apr – jun 2020', file: 'CSR-Report-2020-1st-Quarter-Apr-Jun-2020-W.pdf' },
      { label: '2nd Quarter', period: 'Jul – Sep 2020', file: 'CSR-Report-2020-2nd-Quarter-Jul-Sep-2020-f.pdf' },
      { label: '3rd Quarter', file: null },
      { label: '4th Quarter', file: null },
    ]
  },
]

const openPdf = (filename) => {
  activePdf.value = `/files/${filename}`
  activePdfName.value = filename.replace(/[-_]/g, ' ').replace('.pdf', '')
}
</script>

<style scoped>
/* ─── HERO ─────────────────────────────────────────────── */
.csrr-hero {
  position: relative;
  background: #0d1f17;
  padding: 4rem 0;
  overflow: hidden;
}
.csrr-hero__bg {
  position: absolute;
  inset: 0;
  background-image:
    radial-gradient(circle at 75% 30%, rgba(34,197,94,0.1) 0%, transparent 55%),
    radial-gradient(circle at 10% 70%, rgba(16,185,129,0.07) 0%, transparent 45%);
  pointer-events: none;
}
.csrr-hero__content { position: relative; z-index: 1; }
.csrr-hero__breadcrumb {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 12px;
  color: rgba(255,255,255,0.4);
  letter-spacing: 0.06em;
  text-transform: uppercase;
  margin-bottom: 1.5rem;
}
.csrr-sep { opacity: 0.35; }
.csrr-active { color: #4ade80; }
.csrr-hero__title {
  font-size: clamp(2.5rem, 5vw, 4.5rem);
  font-weight: 900;
  color: #fff;
  letter-spacing: -0.02em;
  margin: 0 0 1rem;
  line-height: 1.05;
}
.csrr-hero__accent { color: #22c55e; }
.csrr-hero__sub {
  font-size: 1rem;
  color: rgba(255,255,255,0.5);
  max-width: 460px;
  line-height: 1.7;
  margin: 0 0 2rem;
}
.csrr-hero__divider {
  width: 52px;
  height: 3px;
  background: #22c55e;
  border-radius: 2px;
}

/* ─── BODY ──────────────────────────────────────────────── */
.csrr-body {
  background: #f8faf8;
  padding: 4rem 0 5rem;
}

/* ─── YEAR BLOCK ─────────────────────────────────────────── */
.csrr-year-block { margin-bottom: 3rem; }
.csrr-year-block:last-child { margin-bottom: 0; }

.csrr-year-header {
  display: flex;
  align-items: center;
  gap: 1rem;
  margin-bottom: 1.25rem;
}
.csrr-year-badge {
  flex-shrink: 0;
  background: #0d1f17;
  color: #4ade80;
  font-size: 0.75rem;
  font-weight: 800;
  letter-spacing: 0.1em;
  padding: 4px 14px;
  border-radius: 20px;
  border: 1px solid rgba(74,222,128,0.25);
}
.csrr-year-line {
  flex-shrink: 0;
  width: 20px;
  height: 2px;
  background: #22c55e;
  border-radius: 2px;
}
.csrr-year-label {
  font-size: 1.1rem;
  font-weight: 800;
  color: #111827;
  letter-spacing: -0.01em;
}

/* ─── QUARTER CARDS ──────────────────────────────────────── */
.csrr-quarters {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 1rem;
}

.csrr-card {
  background: #fff;
  border: 1px solid #e5e7eb;
  border-radius: 14px;
  padding: 1.25rem 1.25rem;
  display: grid;
  grid-template-columns: auto 1fr auto;
  align-items: center;
  gap: 0.875rem;
  min-height: 80px;
  transition: all 0.2s;
}
.csrr-card--available {
  cursor: pointer;
}
.csrr-card--available:hover {
  border-color: #86efac;
  box-shadow: 0 4px 20px rgba(34,197,94,0.1);
  transform: translateY(-2px);
}
.csrr-card--unavailable {
  opacity: 1;
  cursor: default;
}

.csrr-card__icon {
  flex-shrink: 0;
  width: 40px;
  height: 40px;
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
}
.csrr-card--available .csrr-card__icon {
  background: #f0fdf4;
  border: 1px solid #bbf7d0;
  color: #16a34a;
}
.csrr-card--unavailable .csrr-card__icon {
  background: #f3f4f6;
  border: 1px solid #e5e7eb;
  color: #9ca3af;
}

.csrr-card__info { min-width: 0; }
.csrr-card__quarter {
  font-size: 0.875rem;
  font-weight: 700;
  color: #111827;
  margin: 0 0 2px;
  white-space: nowrap;
}
.csrr-card__period {
  font-size: 0.72rem;
  color: #9ca3af;
  margin: 0;
  white-space: nowrap;
}

.csrr-card__action { flex-shrink: 0; }
.csrr-card__btn {
  display: inline-flex;
  align-items: center;
  gap: 5px;
  background: #f0fdf4;
  color: #16a34a;
  border: 1px solid #bbf7d0;
  border-radius: 8px;
  font-size: 0.7rem;
  font-weight: 700;
  padding: 5px 10px;
  white-space: nowrap;
  transition: all 0.15s;
}
.csrr-card--available:hover .csrr-card__btn {
  background: #16a34a;
  color: #fff;
  border-color: #16a34a;
}
.csrr-card__na {
  font-size: 0.68rem;
  color: #d1d5db;
  font-weight: 600;
  white-space: nowrap;
}

/* ─── MODAL ──────────────────────────────────────────────── */
.csrr-modal {
  position: fixed;
  inset: 0;
  background: rgba(0,0,0,0.65);
  z-index: 9999;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 1.5rem;
}
.csrr-modal__box {
  background: #fff;
  border-radius: 16px;
  overflow: hidden;
  width: 100%;
  max-width: 960px;
  height: 88vh;
  display: flex;
  flex-direction: column;
  box-shadow: 0 24px 60px rgba(0,0,0,0.3);
}
.csrr-modal__header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0.875rem 1.25rem;
  border-bottom: 1px solid #e5e7eb;
  background: #f9fafb;
  flex-shrink: 0;
}
.csrr-modal__header-left {
  display: flex;
  align-items: center;
  gap: 0.625rem;
}
.csrr-modal__pdf-icon {
  width: 30px;
  height: 30px;
  background: #f0fdf4;
  border: 1px solid #bbf7d0;
  border-radius: 7px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #16a34a;
  flex-shrink: 0;
}
.csrr-modal__title {
  font-size: 0.8rem;
  font-weight: 700;
  color: #374151;
  text-transform: capitalize;
}
.csrr-modal__close {
  width: 30px;
  height: 30px;
  border-radius: 7px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #6b7280;
  background: transparent;
  border: 1px solid #e5e7eb;
  cursor: pointer;
  transition: all 0.15s;
}
.csrr-modal__close:hover {
  background: #fee2e2;
  color: #dc2626;
  border-color: #fca5a5;
}
.csrr-modal__iframe {
  flex: 1;
  width: 100%;
  border: none;
}

/* ─── FOOTER BAR ─────────────────────────────────────────── */
.csrr-footer-bar {
  height: 5px;
  background: linear-gradient(90deg, #15803d 0%, #22c55e 50%, #86efac 100%);
}

/* ─── RESPONSIVE ─────────────────────────────────────────── */
@media (max-width: 1024px) {
  .csrr-quarters { grid-template-columns: repeat(2, 1fr); }
}
@media (max-width: 640px) {
  .csrr-quarters { grid-template-columns: 1fr; }
  .csrr-modal__box { height: 90vh; }
}
</style>