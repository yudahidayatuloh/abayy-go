package models

// DataProfil → struktur profil sekolah
type DataProfil struct {
    NamaK     string
    Sejarah  string
    Visi     string
    Misi     []string
    Foto     string
}

// IsiProfil → data statis profil sekolah
var IsiProfil = DataProfil{
    NamaK: "Kurniawan S.Pd",
    Sejarah: "Sejalan dengan Program Pemerintah dibidang pendidikan Menengah Kejuruan pada saat itu yakni pemerataan akses ditambah pula dengan banyaknya keinginan masyarakat yang mengharapkan adanya SMK Negeri di daerah Kecamatan Purbaratu Kota Tasikmalaya, terutama untuk menampung tamatan dari SLTP yang ingin melanjutkan ke SMK maka beberapa tokoh masyarakat, unsur pejabat pemerintah di Kecamatan Purbaratu Kota Tasikmalaya mengusulkan kepada pemerintah Kota Tasikmalaya dan Pemerintah Provinsi Jawa Barat, agar berdirinya SMK Negeri di Kecamatan Purbaratu Kota Tasikmalaya. Setelah melalui perjuangan yang sangat panjang dari berbagai pihak khususnya Disdik Kota Tasikmalaya dan pihak-pihak terkait lainnya sesuai prosedur dan ketentuan yang berlaku pada waktu itu dengan mengucapkan syukur Alhamdulillah akhirnya pada tahun pelajaran 2010/2011 SMK Negeri 4 Tasikmalaya mulai berdiri, pada awalnya membuka Kompetensi Keahlian Teknik Komputer dan Jaringan dengan pendaftar peserta Didik Baru pada waktu itu berjumlah 44 orang. Pada awalnya tempat belajar menumpang di SMP Negeri 17 Kota Tasikmalaya, dan sekolah induknya adalah SMK Negeri 2 Kota Tasikmalaya, Untuk Tenaga pendidik dan Tenaga Kependidikan masih dibantu sepenuhnya oleh kedua sekolah tersebut. Pada tahun 2012 keluarlah surat Izin Pendirian berdasarkan keputusan kepala Badan Pelayanan Perizinan Terpadu Kota Tasikmalaya No. 420/9/SK-BPPT/2012 Tanggal 06 Februari 2012.",
    Visi: "Terwujudnya lulusan yang berkarakter, berprestasi, dan berdaya saing tinggi di bidang teknologi" +
        "serta berlandaskan iman dan taqwa.",
    Misi: []string{
        "Menyiapkan Sumber Daya Manusia yang cerdas dan kompeten baik hardskill maupun softskill.",
        "Meningkatkan dan mengembangkan aktifitas serta kreatifitas seluruh warga sekolah dalam berbagai kegiatan positif.",
        "Mewujudkan manajemen pengelolaan yang efektif, efesien, transparan, akuntabel dan layanan prima.",
        "Menumbuhkan dan mengembangkan potensi dan kemampuan Sumber Daya Manusia yang berdaya saing tinggi melalui berbagai kegiatan akademik maupun non akademik.",
        "Menyelenggarakan berbagai program kegiatan dalam upaya meningkatkan Sumber Daya Manusia yang mampu menyesuaikan dengan perkembangan.",
        "Mengembangkan sarana prasarana dan lingkungan sekolah yang menyenangkan sebagai wadah menumbuhkembangkan daya kreasi dan inovasi untuk menghasilkan produk teknologi tepat guna.",
        "Membangun jiwa wirausaha yang handal melalui pembelajaran Teaching Factory (TEFA) dan Kelas Industri.",
        "Menjalin dan mengembangkan kemitraan dengan Industri dan Dunia Kerja (IDUKA) serta lembaga lainnya yang relevan.",
        "Menanamkan dan membudayakan sikap dan perilaku yang baik pada aktivitas di sekolah maupun dalam kehidupan sehari-hari.",
    },
    Foto: "https://smkn4tsm.netlify.app/assets/kepsek.png",
}
