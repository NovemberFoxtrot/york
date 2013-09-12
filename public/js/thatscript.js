var thatApp = angular.module("thatApp", []);

thatApp.controller("thatCtrl", function($scope) {
	$scope.leaveVoicemail = function(number, message) {
		console.log(number, message);
	};
});

thatApp.directive("phone", function() {
	return {
		restrict: "E",
		scope: {
			number:"@",
			network:"=",
			makeCall:"&"
		},
    template: '<div class="panel">Number: {{number}} </div>Network:<select ng-model="network" ng-options="net for net in networks">'+
              '<input type="text" ng-model="value">' +
              '<div class="button" ng-click="makeCall({number: number, message:value})">Call home!</div>',
		link: function(scope) {
			scope.networks = ["Verizon", "AT&T", "Pacific Bell"];
			scope.network = scope.networks[0];
		}
	};
});
