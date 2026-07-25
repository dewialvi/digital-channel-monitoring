(function () {

    const token = localStorage.getItem('token');

    // Jika tidak ada token,
    // redirect ke halaman login

    if (!token) {

        window.location.href = '/login';

        return;

    }


    // Ambil data user dari localStorage

    const user = JSON.parse(
        localStorage.getItem('user') || '{}'
    );


    // Tampilkan nama user di navbar

    const userNameEl =
        document.getElementById('userName');

    if (userNameEl) {

        userNameEl.textContent =
            `${user.name} (${user.role})`;

    }


    // Sembunyikan menu Master User
    // jika user bukan admin

    const menuUsers =
        document.getElementById('menuUsers');

    if (
        menuUsers &&
        user.role !== 'admin'
    ) {

        menuUsers.style.display = 'none';

    }


    // Logout

    const logoutBtn =
        document.getElementById('logoutBtn');

    if (logoutBtn) {

        logoutBtn.addEventListener(
            'click',
            async function () {

                try {

                    await fetch(
                        'http://localhost:8080/api/v1/logout',
                        {
                            method: 'POST',

                            headers: {
                                'Authorization':
                                    `Bearer ${token}`
                            }
                        }
                    );

                } finally {

                    // Hapus token

                    localStorage.removeItem(
                        'token'
                    );

                    // Hapus data user

                    localStorage.removeItem(
                        'user'
                    );

                    // Kembali ke login

                    window.location.href =
                        '/login';

                }

            }
        );

    }

})();