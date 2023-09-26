document.getElementById('searchBar').addEventListener('input', function () {
    let query = this.value.toLowerCase().trim();
    let tableRows = document.querySelectorAll('#artistTable .searchable');
    let suggestions = document.getElementById('suggestions');

    suggestions.innerHTML = ''; // Clear previous suggestions

    if(query === '') {
        tableRows.forEach(row => row.style.display = ''); // Show all rows when the search bar is empty
        return;
    }

    let suggestionsArray = [];

    tableRows.forEach(row => {
        row.style.display = 'none'; // Initially hide all rows

        let cells = row.querySelectorAll('td');
        let artistName = cells[0].textContent.trim().toLowerCase();
        let members = cells[2].textContent.trim().toLowerCase().split(',');

        let creationDate = cells[3].textContent.trim().toLowerCase();
        let firstAlbum = cells[4].textContent.trim().toLowerCase();
        let locations = cells[5].textContent.trim().toLowerCase().split('\n');
        let concertDates = cells[6].textContent.trim().toLowerCase().split('\n');

        if (artistName.includes(query)) {
            suggestionsArray.push({ type: 'artist/band', name: cells[0].textContent.trim() });
            row.style.display = ''; // Show row if artist name matches
        }

        members.forEach(member => {
            let memberName = member.trim();
            if (memberName.includes(query)) {
                suggestionsArray.push({ type: 'member', name: memberName });
                row.style.display = ''; // Show row if member name matches
            }
        });

        if (creationDate.includes(query)) {
            suggestionsArray.push({ type: 'creation date', name: creationDate });
            row.style.display = ''; // Show row if creation date matches
        }

        if (firstAlbum.includes(query)) {
            suggestionsArray.push({ type: 'first album', name: firstAlbum });
            row.style.display = ''; // Show row if first album matches
        }

        locations.forEach(location => {
            let locationName = location.trim();
            if (locationName.includes(query)) {
                suggestionsArray.push({ type: 'location', name: locationName });
                row.style.display = ''; // Show row if location matches
            }
        });

        concertDates.forEach(date => {
            let dateString = date.trim();
            if (dateString.includes(query)) {
                suggestionsArray.push({ type: 'concert date', name: dateString });
                row.style.display = ''; // Show row if concert date matches
            }
        });
    });

    // Display suggestions
    suggestionsArray.forEach(suggestion => {
        let div = document.createElement('div');
        div.textContent = `${suggestion.name} - ${suggestion.type}`;
        suggestions.appendChild(div);
    });
});

document.getElementById('searchBar').addEventListener('input', function() {
    let query = this.value.toLowerCase();
    let tableRows = document.querySelectorAll('#artistTable .searchable');
    let suggestions = document.getElementById('suggestions');
    suggestions.innerHTML = '';

    tableRows.forEach(row => {
        let cellText = row.textContent.toLowerCase();
        if (cellText.includes(query)) {
            row.style.display = '';
            if (query) {
                let suggestionDiv = document.createElement('div');
                suggestionDiv.textContent = row.cells[0].textContent; // Artist Name
                suggestionDiv.addEventListener('click', function() {
                    document.getElementById('searchBar').value = this.textContent;
                    suggestions.innerHTML = '';
                    filterRows(this.textContent.toLowerCase());
                });
                suggestions.appendChild(suggestionDiv);
            }
        } else {
            row.style.display = 'none';
        }
    });
});

function filterRows(query) {
    let tableRows = document.querySelectorAll('#artistTable .searchable');
    tableRows.forEach(row => {
        row.style.display = row.textContent.toLowerCase().includes(query) ? '' : 'none';
    });
}
