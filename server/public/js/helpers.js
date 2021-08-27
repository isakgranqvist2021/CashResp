/** @format */

(function() {
    let closeIcon = document.querySelector('.alert span.material-icons-outlined');

    if(closeIcon !== null) {
        closeIcon.addEventListener('click', () => 
            document.querySelector('.alert').remove());
    }
})(); 