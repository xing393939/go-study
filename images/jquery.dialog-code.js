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
        'newproc', 'newproc1', 'runqput', 'wakep', 'startm', 'mstart', 'releasep', 'acquirep',
        'mainPC', '`type."".Myintinterface`', '`"".(*Myint).fun`', '`"".Myint.fun`', 'schedule',
        '`go.itab."".Myint,"".Myintinterface`', '`go.itab.*"".Myint,"".Myintinterface`',
    ];
    $.get(`../docs/go1.16.10/${word}.html`, function (text) {
        $('#DialogCodeTemp').html(text);
        let newElem = $('#DialogCodeTemp').find('.highlighter-rouge');
        let spans = newElem.find('span');
        spans.each(function (k, span) {
            if (keywords.includes(span.innerText) && word != span.innerText) {
                let subWord = span.innerText.replaceAll('*', '-').replaceAll('"".', '``.')
                spans.eq(k).html(`<a onclick="getSourceCode('${subWord}');">${span.innerText}</a>`)
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