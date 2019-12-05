var endpoint = 'api';

angular.module('rc', ['ngAnimate', 'ngSanitize', 'ui.bootstrap'])
    .controller('extract', ['$scope', '$http', function ($scope, $http) {
        $scope.numLimit = 20;
        $scope.numPerPage = 10;
        $scope.maxSize = 5;
        $scope.informationUrl = informationUrl;

        $scope.extractListFiltered = {};
        $scope.extractList = {};

        $http({method: 'GET', url: endpoint + '/horodatage', headers: {"Access-Control-Allow-Credentials" : true}, withCredentials: true}).then(function (response) {
            var headers = response.headers();
            localStorage.setItem('csrfToken', headers['x-csrf-token']);

            ret = response.data;
            ret.sort(function (a, b) {
                return b.date - a.date;
            });
            $scope.extractList = ret;
            $scope.extractListFiltered = $scope.extractList.slice(0, $scope.numPerPage);
            $scope.totalItems = response.data.length;
        }, function (response) {
            console.log(response);
        });
        $scope.delete_action = function () {
            hashes = $scope.extractListFiltered
                .filter((a) => a.checked)
                .reduce((acc, val) => acc.concat([val.hash]), []);
            $http({method: 'POST', url: endpoint + '/recu', data: hashes, headers: { "X-CSRF-Token": localStorage.getItem("csrfToken"), "Access-Control-Allow-Credentials" : true}, withCredentials: true}).then(function (response) {
                location.reload();
                console.log(response);
            }, function (response) {
                console.log(response);
            });
        };

        $scope.pageChanged = function () {
            var begin = (($scope.currentPage - 1) * $scope.numPerPage);
            var end = begin + $scope.numPerPage;

            $scope.extractListFiltered = $scope.extractList.slice(begin, end);
        };
    }]);

function getParameterByName(name, url) {
    if (!url) {
        url = window.location.href;
    }
    name = name.replace(/[\[\]]/g, "\\$&");
    var regex = new RegExp("[?&]" + name + "(=([^&#]*)|&|#|$)"),
        results = regex.exec(url);
    if (!results) return null;
    if (!results[2]) return '';
    return decodeURIComponent(results[2].replace(/\+/g, " "));
}

Dropzone.autoDiscover = false;
var uploadextraitzone;

function successmultiple(files, message, e) {
    from = message.from;
    target_hash = message.target_hash;
    t = new Date(message.time * 1000);
    $("#infobox").html(display_str);
    $("#infobox").attr("class", "alert alert-success");
    setTimeout(function () {
        location.reload();
    }, 25000); //TODO wait by polling+error handling
}

function errormultiple(files, message, e) {
    $("#infobox").text(message);
    $("#infobox").attr("class", "alert alert-danger");
    console.log(message);
}

$(function () {
    uploadextraitzone = new Dropzone("div#uploadextraitzone", {
        url: endpoint + "/upload",
        headers: {'Access-Control-Allow-Credentials' : true},
        withCredentials: true,
        uploadMultiple: true,
        paramName: "myfiles",
        dictDefaultMessage: dictDefaultMessage,
        dictFallbackMessage: dictDefaultMessage,
        parallelUploads: 256,
        autoProcessQueue: false,
        successmultiple: successmultiple,
        errormultiple: errormultiple,
        addRemoveLinks: true,
    });
});

function processValidate() {
    uploadextraitzone.options.headers['X-CSRF-Token'] = localStorage.getItem("csrfToken");
    uploadextraitzone.processQueue();
}
