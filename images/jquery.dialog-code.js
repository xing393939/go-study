var gitbook = gitbook || [];
gitbook.push(function() {
    $('body').append('<div id="DialogCodeTemp" style="display:none"></div>');
    $.getScript("../images/jquery.dialog.js", function () {
        document.querySelectorAll('.DialogCode').forEach(el => {
            getSourceCode($(el).data('code'), el)
        })
    });
});

function getSourceCode(word, parent) {
    let keywords = [
        'newproc', 'newproc1', 'runqput', 'wakep', 'startm',
        'mainPC', '`type.``.Myintinterface`', '```.(-Myint).fun`',
    ];
    $.get(`../docs/go1.16.10/${word}.html`, function (text) {
        $('#DialogCodeTemp').html(text.replaceAll('(*', '(-').replaceAll('"".', '``.'));
        let newElem = $('#DialogCodeTemp').find('.highlighter-rouge');
        let spans = newElem.find('span');
        spans.each(function (k, span) {
            if (keywords.includes(span.innerText) && word != span.innerText) {
                spans.eq(k).html(`<a onclick="getSourceCode('${span.innerText}');">${span.innerText}</a>`)
            }
        });
        if (parent) {
            $(parent).append(newElem)
        } else {
            $('body').append(newElem);
            newElem.dialog({
                'title': word
            }, function (api) {
                api.open();
            });
        }
    })
}