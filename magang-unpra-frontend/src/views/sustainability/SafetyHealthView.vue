<template>

  <!-- ── HERO ── -->
  <section class="relative h-[400px] flex items-end bg-cover bg-center overflow-hidden"
    style="background-image: url('https://images.unsplash.com/photo-1504307651254-35680f356dfd?w=1920')">
    <div class="absolute inset-0 bg-gradient-to-t from-black/85 via-black/45 to-transparent"></div>

    <div class="absolute top-10 left-10 flex items-center gap-2">
      <span class="w-1 h-5 bg-green-400 rounded-full"></span>
      <span class="text-green-400 text-xs font-bold tracking-[3px] uppercase">Sustainability</span>
    </div>

    <div class="relative z-10 max-w-5xl mx-auto px-10 pb-14 w-full">
      <h1 class="text-5xl font-black text-white leading-tight mb-4">
        Safety & Health <span class="text-green-400">Office</span>
      </h1>
      <p class="text-gray-300 text-sm max-w-xl leading-relaxed">
        Committed to implementing Occupational Safety and Health across all factory areas for every employee, contractor and supplier.
      </p>
      <div class="flex items-center gap-2 mt-5 text-xs text-gray-400">
        <span class="hover:text-white cursor-pointer transition">Home</span>
        <span class="text-gray-600">/</span>
        <span class="hover:text-white cursor-pointer transition">Sustainability</span>
        <span class="text-gray-600">/</span>
        <span class="text-green-400 font-semibold">Safety & Health Office</span>
      </div>
    </div>
  </section>

  <!-- ── STAT BAR ── -->
  <section class="bg-green-700 py-5">
    <div class="max-w-5xl mx-auto px-10">
      <div class="grid grid-cols-3 divide-x divide-green-600">
        <div class="text-center px-6">
          <p class="text-xl font-black text-white">SMK3 <span class="text-green-300 text-sm font-semibold">& OHSAS 18001</span></p>
          <p class="text-green-200 text-xs tracking-widest uppercase mt-1">Certified System</p>
        </div>
        <div class="text-center px-6">
          <p class="text-2xl font-black text-white">0% <span class="text-green-300 text-base font-semibold">Target</span></p>
          <p class="text-green-200 text-xs tracking-widest uppercase mt-1">Occupational Diseases (PAK)</p>
        </div>
        <div class="text-center px-6">
          <p class="text-2xl font-black text-white">10x <span class="text-green-300 text-base font-semibold">/ Year</span></p>
          <p class="text-green-200 text-xs tracking-widest uppercase mt-1">Safety Training & Patrol</p>
        </div>
      </div>
    </div>
  </section>

  <!-- ── K3 POLICY ── -->
  <section class="py-20 bg-white">
    <div class="max-w-5xl mx-auto px-10">

      <div class="flex items-center gap-4 mb-8">
        <div class="w-12 h-12 rounded-2xl bg-green-600 flex items-center justify-center text-2xl flex-shrink-0">
          🛡️
        </div>
        <div>
          <span class="inline-block bg-green-100 text-green-700 text-xs font-bold tracking-widest uppercase px-3 py-1 rounded-full mb-1">
            K3 Policy
          </span>
          <h2 class="text-3xl font-black text-gray-900">Kebijakan K3</h2>
        </div>
      </div>

      <div class="grid grid-cols-5 gap-10">

        <!-- Teks -->
        <div class="col-span-3 space-y-4 text-gray-500 text-sm leading-relaxed">
          <p>Tanjungenim Lestari Pulp and Paper is fully committed to implementing Occupational Safety and Health in all factory areas by each of our employees, contractors and suppliers.</p>
          <p>The company believes that all employees are valuable assets for the company. This policy is a commitment in the field of Occupational Safety and Health to support all company activities with high productivity without creating risks to Occupational Safety and Health for employees and related parties.</p>
          <p>To ensure the efficient implementation of the management system, we have implemented an Occupational Safety and Health management system – <strong class="text-gray-700">SMK3 and OHSAS 18001</strong> with a spirit of continuous improvement and compliance with government regulations, the Company's Occupational Safety and Health Policy and other safety requirements.</p>
          <p>The company's attention is realized by providing training, programs and all the resources needed to ensure compliance with Occupational Safety and Health requirements.</p>
        </div>

        <!-- Highlight cards -->
        <div class="col-span-2 space-y-3">
          <div v-for="card in policyCards" :key="card.title"
            class="bg-gray-50 border border-gray-100 rounded-2xl p-5 hover:border-green-200 hover:bg-green-50/50 transition">
            <div class="text-xl mb-2">{{ card.icon }}</div>
            <p class="text-gray-900 font-bold text-sm mb-1">{{ card.title }}</p>
            <p class="text-gray-400 text-xs leading-relaxed">{{ card.desc }}</p>
          </div>
        </div>
      </div>

      <!-- Bottom note -->
      <div class="mt-8 border-l-4 border-green-500 pl-5 bg-green-50 py-4 pr-5 rounded-r-xl">
        <p class="text-green-800 text-sm font-medium leading-relaxed">
          ⚠️ Occupational Safety and Health Policy of PT. Tanjungenim Lestari Pulp and Paper must be obeyed and followed by every employee, including contractors and suppliers seriously.
        </p>
      </div>
    </div>
  </section>

  <!-- ── KOMITE K3 ── -->
  <section class="py-20 bg-gray-50">
    <div class="max-w-5xl mx-auto px-10">

      <div class="flex items-center gap-4 mb-10">
        <div class="w-12 h-12 rounded-2xl bg-green-600 flex items-center justify-center text-2xl flex-shrink-0">
          👥
        </div>
        <div>
          <span class="inline-block bg-green-100 text-green-700 text-xs font-bold tracking-widest uppercase px-3 py-1 rounded-full mb-1">
            Komite K3
          </span>
          <h2 class="text-3xl font-black text-gray-900">P2K3 Committee</h2>
        </div>
      </div>

      <div class="grid grid-cols-2 gap-6">
        <div class="bg-white rounded-2xl p-6 border border-gray-100 shadow-sm hover:shadow-md transition">
          <div class="w-9 h-9 rounded-xl bg-green-100 flex items-center justify-center text-base mb-4">🏢</div>
          <h3 class="font-bold text-gray-900 text-sm mb-2">Safety Office & P2K3</h3>
          <p class="text-gray-400 text-sm leading-relaxed">
            PT. Tanjungenim Lestari Pulp and Paper has a special section that accommodates the implementation of the K3 program, namely the Safety Office and the K3 Committee or the Occupational Safety and Health Management Committee (P2K3) consisting of management lines and employee representatives in each Department / Section.
          </p>
        </div>
        <div class="bg-white rounded-2xl p-6 border border-gray-100 shadow-sm hover:shadow-md transition">
          <div class="w-9 h-9 rounded-xl bg-green-100 flex items-center justify-center text-base mb-4">📋</div>
          <h3 class="font-bold text-gray-900 text-sm mb-2">Legal Authority</h3>
          <p class="text-gray-400 text-sm leading-relaxed">
            The tasks and authority of P2K3 have been decided in the Decree of the Head of the South Sumatra Province Manpower and Transmigration Office No.254 / SK / Nakertrans / 2016. P2K3 in accordance with regulations of the Minister of Manpower has the duty to provide an output of consideration whether requested or not requested by management regarding aspects of occupational safety and health.
          </p>
        </div>
      </div>
    </div>
  </section>

  <!-- ── K3 TARGET ── -->
  <section class="py-20 bg-white">
    <div class="max-w-5xl mx-auto px-10">

      <div class="flex items-center gap-4 mb-4">
        <div class="w-12 h-12 rounded-2xl bg-green-600 flex items-center justify-center text-2xl flex-shrink-0">
          🎯
        </div>
        <div>
          <span class="inline-block bg-green-100 text-green-700 text-xs font-bold tracking-widest uppercase px-3 py-1 rounded-full mb-1">
            Target & Objective
          </span>
          <h2 class="text-3xl font-black text-gray-900">K3 Target and Objective</h2>
        </div>
      </div>

      <p class="text-gray-400 text-sm mb-10 ml-16">
        Key performance targets set to ensure occupational safety and health excellence across all operations.
      </p>

      <div class="space-y-3">
        <div v-for="(target, i) in k3Targets" :key="i"
          class="flex gap-4 bg-gray-50 border border-gray-100 rounded-xl p-5 hover:border-green-200 hover:bg-green-50/40 transition group">
          <div class="w-9 h-9 rounded-xl bg-green-600 text-white font-black text-xs flex items-center justify-center flex-shrink-0 group-hover:bg-green-700 transition">
            {{ String(i + 1).padStart(2, '0') }}
          </div>
          <p class="text-gray-600 text-sm leading-relaxed self-center">{{ target }}</p>
        </div>
      </div>

      <!-- Reduction target banner -->
      <div class="mt-8 bg-green-700 rounded-2xl p-6 text-center">
        <p class="text-green-200 text-xs font-bold tracking-widest uppercase mb-2">Annual Reduction Target</p>
        <div class="grid grid-cols-3 gap-6 mt-4">
          <div>
            <p class="text-white text-3xl font-black">-10%</p>
            <p class="text-green-300 text-xs mt-1">Occupational Accidents (vs 2018)</p>
          </div>
          <div class="border-x border-green-600">
            <p class="text-white text-3xl font-black">-5%</p>
            <p class="text-green-300 text-xs mt-1">Industrial Fires</p>
          </div>
          <div>
            <p class="text-white text-3xl font-black">0%</p>
            <p class="text-green-300 text-xs mt-1">Occupational Diseases (PAK)</p>
          </div>
        </div>
      </div>
    </div>
  </section>

  <!-- ── PROGRAM K3 ── -->
  <section class="py-20 bg-gray-50">
    <div class="max-w-5xl mx-auto px-10">

      <div class="flex items-center gap-4 mb-4">
        <div class="w-12 h-12 rounded-2xl bg-green-600 flex items-center justify-center text-2xl flex-shrink-0">
          ⚙️
        </div>
        <div>
          <span class="inline-block bg-green-100 text-green-700 text-xs font-bold tracking-widest uppercase px-3 py-1 rounded-full mb-1">
            Program K3
          </span>
          <h2 class="text-3xl font-black text-gray-900">K3 Programs</h2>
        </div>
      </div>

      <p class="text-gray-500 text-sm leading-relaxed mb-10 ml-16 max-w-3xl">
        In achieving continuous improvement and ensuring that the system is well implemented, PT. Tanjungenim Lestari Pulp and Paper organizes K3 efforts that are contained in the K3 program. K3 program activities become a program that is integrated with the company's operational processes so that there is a mutually beneficial synergy without reducing one part. Some of the K3 programs carried out in 2019 are:
      </p>

      <div class="grid grid-cols-2 gap-4">
        <div v-for="(program, i) in k3Programs" :key="i"
          class="flex gap-4 bg-white border border-gray-100 rounded-xl p-5 hover:shadow-md hover:border-green-100 transition group">
          <div class="w-8 h-8 rounded-lg bg-green-600 text-white font-black text-xs flex items-center justify-center flex-shrink-0 group-hover:bg-green-700 transition">
            {{ String(i + 1).padStart(2, '0') }}
          </div>
          <p class="text-gray-600 text-sm leading-relaxed self-center">{{ program }}</p>
        </div>
      </div>

      <!-- Foto Program -->
      <div class="mt-12 rounded-2xl overflow-hidden shadow-xl aspect-video">
        <img
          src="https://images.unsplash.com/photo-1504307651254-35680f356dfd?w=1200"
          alt="K3 Training Program"
          class="w-full h-full object-cover hover:scale-105 transition-transform duration-700"
        />
      </div>
      <p class="text-center text-gray-400 text-xs mt-3 italic">
        K3 Training & Emergency Response Program — PT. Tanjungenim Lestari Pulp and Paper
      </p>

    </div>
  </section>

  <!-- ── CTA ── -->
  <section class="py-16 bg-green-700">
    <div class="max-w-5xl mx-auto px-10 text-center">
      <h2 class="text-3xl font-black text-white mb-3">
        Safety is Everyone's Responsibility
      </h2>
      <p class="text-green-200 text-sm mb-8 max-w-lg mx-auto">
        TEL is committed to maintaining zero accidents through continuous training, strict compliance, and a strong safety culture.
      </p>

    </div>
  </section>

</template>

<script setup>
const policyCards = [
  {
    icon: '📋',
    title: 'SMK3 & OHSAS 18001',
    desc: 'Certified management system ensuring continuous improvement and full compliance with government safety regulations.'
  },
  {
    icon: '🎓',
    title: 'Training & Resources',
    desc: 'All resources needed to ensure full compliance with Occupational Safety and Health requirements are provided.'
  },
  {
    icon: '🤝',
    title: 'All Parties Covered',
    desc: 'Policy covers all employees, contractors and suppliers operating within TEL factory areas.'
  },
]

const k3Targets = [
  'Update and adopt government regulations related to OSH by 100% in 2019.',
  'Renewed operating license deeds for all machines and OSH-related licenses for all equipment (lifts, lifting equipment, boilers and electricity) by 96% in 2019 issued by the Ministry of Manpower of the Republic of Indonesia.',
  'Ensure that companies, contractors, suppliers and vendors comply with regulations by 100% in 2019.',
  'Increase awareness of safety and safety in OHS risk assessments by providing training at least 10 times a year.',
  'Implement safety patrols a minimum of 10 times in one year.',
  'Conduct a hazard assessment in the work environment at least once a year.',
]

const k3Programs = [
  'Comply with government regulations related to occupational safety and health, both new and amended old regulations.',
  'Equip all work equipment and operational machinery with an operational permit.',
  'Ensuring companies, supplier contractors and vendors meet SMK3 requirements.',
  'Perform medical checkups.',
  'Conduct Training Need Analysis (TNA) related to K3 and conduct training for employees.',
  'K3 Promotion.',
  'Patrol and follow up findings / finding / K3 abnormality.',
  'Measure the Work Environment once a year and follow up abnormal parameters.',
  'Perform preventive maintenance for fire equipment.',
  'Implement work permits and fire work permits in the work environment.',
  'Investigating occupational accidents, Occupational Diseases (PAK) and fires.',
]
</script>