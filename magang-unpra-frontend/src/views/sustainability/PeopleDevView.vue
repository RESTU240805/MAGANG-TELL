<template>

  <!-- ── HERO ── -->
  <section class="relative h-[420px] flex items-end bg-cover bg-center overflow-hidden"
    style="background-image: url('/images/people3.jpeg')">
    <div class="absolute inset-0 bg-gradient-to-t from-black/80 via-black/40 to-transparent"></div>

    <div class="absolute top-10 left-10 flex items-center gap-2 opacity-0 anim-fade-right" style="animation-delay: 0.1s">
      <span class="w-1 h-5 bg-green-400 rounded-full"></span>
      <span class="text-green-400 text-xs font-bold tracking-[3px] uppercase">Sustainability</span>
    </div>

    <div class="relative z-10 max-w-5xl mx-auto px-10 pb-14 w-full">
      <h1 class="text-5xl font-black text-white leading-tight mb-4 opacity-0 anim-fade-up" style="animation-delay: 0.2s">
        People <span class="text-green-400">Development</span>
      </h1>
      <p class="text-gray-300 text-sm max-w-xl leading-relaxed opacity-0 anim-fade-up" style="animation-delay: 0.3s">
        Preparing the best talent to achieve the highest performance and business sustainability.
      </p>
      <div class="flex items-center gap-2 mt-5 text-xs text-gray-400 opacity-0 anim-fade-up" style="animation-delay: 0.4s">
        <RouterLink to="/" class="hover:text-white cursor-pointer transition">Home</RouterLink>
        <span class="text-gray-600">/</span>
        <span class="hover:text-white cursor-pointer transition">Sustainability</span>
        <span class="text-gray-600">/</span>
        <span class="text-green-400 font-semibold">People Development</span>
      </div>
    </div>
  </section>

  <!-- ── PHOTO GRID ── -->
  <section class="py-16 bg-white">
    <div class="max-w-5xl mx-auto px-10">
      <div class="grid grid-cols-3 gap-3 h-64 rounded-2xl overflow-hidden">
        <div class="col-span-1 overflow-hidden anim-item">
          <img src="/images/Training-First-Aid.jpg"
            alt="Training" class="w-full h-full object-cover hover:scale-105 transition-transform duration-500"/>
        </div>
        <div class="col-span-1 overflow-hidden anim-item" style="animation-delay: 0.1s">
          <img src="/images/Management-Safety-Patrol.jpg"
            alt="Workshop" class="w-full h-full object-cover hover:scale-105 transition-transform duration-500"/>
        </div>
        <div class="col-span-1 overflow-hidden anim-item" style="animation-delay: 0.2s">
          <img src="/images/Industrial-Hygiene.jpg"
            alt="Team" class="w-full h-full object-cover hover:scale-105 transition-transform duration-500"/>
        </div>
      </div>
    </div>
  </section>

  <!-- ── INTRO TEXT ── -->
  <section class="pb-16 bg-white">
    <div class="max-w-5xl mx-auto px-10">
      <div class="grid grid-cols-5 gap-12">

        <div class="col-span-3 space-y-4 text-gray-500 text-sm leading-relaxed anim-slide-right">
          <p>
            People development is carried out in order to prepare the best talent to be able to show the
            best performance and to achieve the business targets they carry. Talking about business is not
            only about running but also how they manage the business even more and sustain it in the future.
          </p>
          <p>
            PT Tanjungenim Lestari Pulp and Paper are committed to developing competent, motivated and
            integrated information systems to achieve a competitive organization. The commitment is realized
            in several activities and achievements as follows:
          </p>
        </div>

        <div class="col-span-2 bg-green-50 border border-green-100 rounded-2xl p-6 flex flex-col justify-center anim-slide-left">
          <div class="w-10 h-10 rounded-xl bg-green-100 flex items-center justify-center mb-4">
            <span class="text-green-600 text-lg">🌱</span>
          </div>
          <p class="text-green-800 font-bold text-base leading-snug mb-2">
            Competent. Motivated. Integrated.
          </p>
          <p class="text-green-600 text-xs leading-relaxed">
            Building people who are ready to sustain and grow the business for the long term.
          </p>
        </div>

      </div>
    </div>
  </section>

  <!-- ── COMMITMENT LIST ── -->
  <section class="py-16 bg-[#14532d]">
    <div class="max-w-5xl mx-auto px-10">

      <div class="text-center mb-12">
        <span class="inline-block bg-green-800 text-green-200 text-xs font-bold tracking-widest uppercase px-3 py-1 rounded-full mb-3 anim-item">
          Our Commitments
        </span>
        <h2 class="text-3xl font-black text-white anim-item">Development Pillars</h2>
        <p class="text-green-200/60 text-sm mt-2 max-w-md mx-auto anim-item">
          Five core programs that shape TEL's approach to people development.
        </p>
      </div>

      <div class="space-y-4">
        <div v-for="(item, index) in commitments" :key="index"
          class="flex gap-5 bg-white/10 backdrop-blur-sm border border-white/10 rounded-2xl p-6 hover:bg-white/15 transition group anim-item"
          :style="{ animationDelay: (index * 0.1) + 's' }">

          <div class="w-10 h-10 rounded-xl bg-green-500 text-white font-black text-sm flex items-center justify-center flex-shrink-0 group-hover:bg-green-400 transition">
            {{ String(index + 1).padStart(2, '0') }}
          </div>

          <div>
            <h3 class="text-white font-bold text-sm mb-1">{{ item.title }}</h3>
            <p class="text-green-200/60 text-sm leading-relaxed">{{ item.desc }}</p>
          </div>
        </div>
      </div>

    </div>
  </section>

  <!-- ── IMAGE SLIDER ── -->
  <section class="py-20 bg-white">
    <div class="max-w-5xl mx-auto px-10">

      <div class="text-center mb-10">
        <span class="inline-block bg-green-100 text-green-700 text-xs font-bold tracking-widest uppercase px-3 py-1 rounded-full mb-3 anim-item">
          Gallery
        </span>
        <h2 class="text-3xl font-black text-gray-900 anim-item">People in Action</h2>
      </div>

      <!-- Slider wrapper -->
      <div class="relative rounded-2xl overflow-hidden shadow-xl aspect-video bg-gray-100 anim-item">

        <img
          v-for="(slide, i) in slides"
          :key="i"
          :src="slide.url"
          :alt="slide.caption"
          class="absolute inset-0 w-full h-full object-cover"
          :style="{ opacity: activeSlide === i ? 1 : 0, transition: 'opacity 0.8s ease-in-out' }"
        />

        <!-- Caption -->
        <div class="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/70 to-transparent p-6">
          <p class="text-white text-sm font-semibold">{{ slides[activeSlide].caption }}</p>
          <p class="text-gray-300 text-xs mt-1">{{ activeSlide + 1 }} / {{ slides.length }}</p>
        </div>

        <!-- Prev -->
        <button @click="manualNav(prev)"
          class="absolute left-4 top-1/2 -translate-y-1/2 w-10 h-10 rounded-full bg-black/50 text-white flex items-center justify-center hover:bg-black/70 transition backdrop-blur-sm">
          ‹
        </button>

        <!-- Next -->
        <button @click="manualNav(next)"
          class="absolute right-4 top-1/2 -translate-y-1/2 w-10 h-10 rounded-full bg-black/50 text-white flex items-center justify-center hover:bg-black/70 transition backdrop-blur-sm">
          ›
        </button>
      </div>

      <!-- Dot Pagination -->
      <div class="flex justify-center gap-2 mt-6">
        <button
          v-for="(_, i) in slides" :key="i"
          @click="manualNav(() => activeSlide = i)"
          class="transition-all duration-300 rounded-full cursor-pointer"
          :class="activeSlide === i ? 'bg-green-600 w-6 h-2.5' : 'bg-gray-200 hover:bg-gray-300 w-2.5 h-2.5'">
        </button>
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
import { ref, onMounted, onUnmounted, nextTick } from 'vue'
import { RouterLink } from 'vue-router'

const activeSlide = ref(0)
let sliderTimer = null

const commitments = [
  {
    title: 'Skills-Based Leadership Coaching',
    desc: 'Achieve a high performing organization by developing skills-based Coaching for leadership styles through a comprehensive Supervisory / Management / Leadership Program.'
  },
  {
    title: 'Succession Planning & Career Path',
    desc: 'Develop and implement succession plans for critical positions to ensure the regeneration of employees through the Career Path and Talent Management.'
  },
  {
    title: 'Community & Scholarship Programs',
    desc: 'Prepare a generation of quality sources from internal employees as well as from the village community through scholarships, internship, community education, school foundation programs and infrastructure improvement.'
  },
  {
    title: 'Technology & Information Skills',
    desc: 'Improve knowledge and skills related to technology and information across all levels of the organization.'
  },
  {
    title: 'Rules, Competencies & Knowledge',
    desc: 'Meet the rules and competencies with good knowledge and skills to maintain professional standards throughout the company.'
  },
]

const slides = [
  { url: '/images/Training-First-Aid.jpg', caption: 'First Aid Training Session' },
  { url: '/images/Management-Safety-Patrol.jpg', caption: 'Safety Patrol Management' },
  { url: '/images/Industrial-Hygiene.jpg', caption: 'Team Safety Workshop' },
  { url: '/images/people1.jpeg', caption: 'Industrial Hygiene Program' },
  { url: '/images/people2.jpeg', caption: 'People Development Initiative' },
]

const next = () => { activeSlide.value = (activeSlide.value + 1) % slides.length }
const prev = () => { activeSlide.value = (activeSlide.value - 1 + slides.length) % slides.length }
const startSlider = () => { sliderTimer = setInterval(next, 5000) }
const stopSlider = () => { clearInterval(sliderTimer) }
const manualNav = (fn) => { stopSlider(); fn(); startSlider() }

// ─── Scroll Observer ──────────────────────────────────────────
let observer = null

onMounted(async () => {
  await nextTick()
  initObserver()
  startSlider()
})

onUnmounted(() => {
  if (observer) {observer.disconnect()}
  stopSlider()
})

const initObserver = () => {
  observer = new IntersectionObserver((entries) => {
    entries.forEach(entry => {
      if (entry.isIntersecting) {
        entry.target.classList.add('anim-visible')
        observer.unobserve(entry.target)
      }
    })
  }, { threshold: 0.15 })

  document.querySelectorAll('.anim-item, .anim-slide-right, .anim-slide-left').forEach(el => {
    observer.observe(el)
  })
}
</script>

<style scoped>
/* Hero animations */
@keyframes fadeUp {
  from { opacity: 0; transform: translateY(24px); }
  to { opacity: 1; transform: translateY(0); }
}
.anim-fade-up {
  animation: fadeUp 0.8s cubic-bezier(0.22, 1, 0.36, 1) forwards;
}

@keyframes fadeRight {
  from { opacity: 0; transform: translateX(-24px); }
  to { opacity: 1; transform: translateX(0); }
}
.anim-fade-right {
  animation: fadeRight 0.8s cubic-bezier(0.22, 1, 0.36, 1) forwards;
}

/* Scroll animations */
.anim-item {
  opacity: 0;
  transform: translateY(30px);
  transition: opacity 0.7s cubic-bezier(0.22, 1, 0.36, 1), transform 0.7s cubic-bezier(0.22, 1, 0.36, 1);
}
.anim-item.anim-visible {
  opacity: 1;
  transform: translateY(0);
}

.anim-slide-right {
  opacity: 0;
  transform: translateX(-40px);
  transition: opacity 0.7s cubic-bezier(0.22, 1, 0.36, 1), transform 0.7s cubic-bezier(0.22, 1, 0.36, 1);
}
.anim-slide-right.anim-visible {
  opacity: 1;
  transform: translateX(0);
}

.anim-slide-left {
  opacity: 0;
  transform: translateX(40px);
  transition: opacity 0.7s cubic-bezier(0.22, 1, 0.36, 1), transform 0.7s cubic-bezier(0.22, 1, 0.36, 1);
}
.anim-slide-left.anim-visible {
  opacity: 1;
  transform: translateX(0);
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
