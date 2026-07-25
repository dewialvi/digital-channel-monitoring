const API_BASE_URL = 'http://localhost:8080/api/v1';

const token = localStorage.getItem('token');

let currentPage = 1;


// =========================
// BADGE SEVERITY
// =========================

const severityBadge = {
    critical: 'bg-danger',
    high: 'bg-danger',
    medium: 'bg-warning text-dark',
    low: 'bg-secondary'
};


// =========================
// LOAD BUG REPORTS
// =========================

async function loadBugReports(page = 1) {

    currentPage = page;

    const search = document.getElementById('searchInput').value;
    const severity = document.getElementById('filterSeverity').value;
    const status = document.getElementById('filterStatus').value;

    const params = new URLSearchParams({
        page: page,
        limit: 10,
        search: search,
        severity: severity,
        status: status
    });

    try {

        const response = await fetch(
            `${API_BASE_URL}/bug-reports?${params}`,
            {
                method: 'GET',
                headers: {
                    'Authorization': `Bearer ${token}`
                }
            }
        );


        const result = await response.json();


        if (!response.ok) {

            throw new Error(
                result.message || 'Gagal memuat data bug reports'
            );

        }


        console.log('Bug Reports API Response:', result);


        renderTable(result.data);

        renderPagination(result.pagination);


    } catch (error) {

        console.error('Error:', error);

        document.getElementById('bugTableBody').innerHTML = `
            <tr>
                <td
                    colspan="6"
                    class="text-center text-danger"
                >
                    ${escapeHtml(error.message)}
                </td>
            </tr>
        `;

    }

}


// =========================
// RENDER TABLE
// =========================

function renderTable(bugs) {

    const tbody = document.getElementById('bugTableBody');


    if (!bugs || bugs.length === 0) {

        tbody.innerHTML = `
            <tr>
                <td
                    colspan="6"
                    class="text-center"
                >
                    Tidak ada data bug report
                </td>
            </tr>
        `;

        return;
    }


    tbody.innerHTML = bugs.map(bug => `

        <tr>

            <td>
                ${escapeHtml(bug.title)}
            </td>


            <td>

                <span
                    class="badge ${severityBadge[bug.severity] || 'bg-secondary'}"
                >
                    ${escapeHtml(bug.severity)}
                </span>

            </td>


            <td>
                ${escapeHtml(bug.priority)}
            </td>


            <td>

                <span class="badge bg-secondary">
                    ${escapeHtml(bug.status)}
                </span>

            </td>


            <td>

                ${
                    bug.reporter
                        ? escapeHtml(bug.reporter.name)
                        : '-'
                }

            </td>


            <td>
                ${formatDate(bug.created_at)}
            </td>

        </tr>

    `).join('');

}


// =========================
// PAGINATION
// =========================

function renderPagination(pagination) {

    const paginationElement =
        document.getElementById('pagination');


    if (!pagination || pagination.total_pages <= 1) {

        paginationElement.innerHTML = '';

        return;
    }


    let html = '';


    // Previous

    html += `

        <li
            class="page-item ${pagination.page === 1 ? 'disabled' : ''}"
        >

            <a
                class="page-link"
                href="#"
                onclick="
                    loadBugReports(${pagination.page - 1});
                    return false;
                "
            >
                Previous
            </a>

        </li>

    `;


    // Number halaman

    for (
        let i = 1;
        i <= pagination.total_pages;
        i++
    ) {

        html += `

            <li
                class="page-item ${i === pagination.page ? 'active' : ''}"
            >

                <a
                    class="page-link"
                    href="#"
                    onclick="
                        loadBugReports(${i});
                        return false;
                    "
                >
                    ${i}
                </a>

            </li>

        `;

    }


    // Next

    html += `

        <li
            class="page-item ${
                pagination.page === pagination.total_pages
                    ? 'disabled'
                    : ''
            }"
        >

            <a
                class="page-link"
                href="#"
                onclick="
                    loadBugReports(${pagination.page + 1});
                    return false;
                "
            >
                Next
            </a>

        </li>

    `;


    paginationElement.innerHTML = html;

}


// =========================
// FORMAT DATE
// =========================

function formatDate(dateString) {

    if (!dateString) {
        return '-';
    }


    return new Date(dateString).toLocaleDateString(
        'id-ID',
        {
            day: '2-digit',
            month: '2-digit',
            year: 'numeric'
        }
    );

}


// =========================
// ESCAPE HTML
// =========================

function escapeHtml(text) {

    if (text === null || text === undefined) {
        return '';
    }


    const div = document.createElement('div');

    div.textContent = text;

    return div.innerHTML;

}


// =========================
// FILTER BUTTON
// =========================

document
    .getElementById('btnFilter')
    .addEventListener(
        'click',
        function () {

            loadBugReports(1);

        }
    );


// =========================
// SEARCH DENGAN ENTER
// =========================

document
    .getElementById('searchInput')
    .addEventListener(
        'keydown',
        function (event) {

            if (event.key === 'Enter') {

                loadBugReports(1);

            }

        }
    );


// =========================
// INITIAL LOAD
// =========================

loadBugReports();