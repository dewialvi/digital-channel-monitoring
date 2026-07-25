const API_BASE_URL =
    'http://localhost:8080/api/v1';

const token =
    localStorage.getItem('token');


async function loadStats() {

    try {

        const response = await fetch(
            `${API_BASE_URL}/bug-reports?limit=100`,
            {
                headers: {
                    'Authorization':
                        `Bearer ${token}`
                }
            }
        );


        if (!response.ok) {

            throw new Error(
                'Gagal memuat data bug reports'
            );

        }


        const result =
            await response.json();


        const bugs =
            result.data || [];


        // =========================
        // HITUNG BUG CRITICAL / HIGH
        // =========================

        const critical =
            bugs.filter(
                bug =>
                    bug.severity === 'critical' ||
                    bug.severity === 'high'
            ).length;


        // =========================
        // HITUNG BUG OPEN
        // =========================

        const open =
            bugs.filter(
                bug =>
                    ![
                        'closed',
                        'verified'
                    ].includes(bug.status)
            ).length;


        // =========================
        // HITUNG BUG CLOSED
        // =========================

        const closed =
            bugs.filter(
                bug =>
                    bug.status === 'closed'
            ).length;


        // =========================
        // TAMPILKAN KE HTML
        // =========================

        document
            .getElementById(
                'statCriticalBugs'
            )
            .textContent = critical;


        document
            .getElementById(
                'statOpenBugs'
            )
            .textContent = open;


        document
            .getElementById(
                'statClosedBugs'
            )
            .textContent = closed;


        // Untuk sementara
        // Feedback belum dibuat API-nya

        document
            .getElementById(
                'statFeedback'
            )
            .textContent = '-';


    } catch (error) {

        console.error(
            'Error loading stats:',
            error
        );

    }

}


loadStats();