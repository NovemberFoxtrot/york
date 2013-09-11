var thatApp = angular.module("thatApp", []);

thatApp.controller("thatCtrl", function($scope) {
	$scope.leaveVoiceMessage = function(number, message) {
		console.log(number, message)
	}
})

thatApp.directive("phone", function() {
	return {
		restrict: "E",
		scope: {
			number:   "@",
			network:  "=",
			makeCall: "&",
		},
		template : '<div>{{number}}</div>' +
		'<input type="text" ng-model="value">' +
    '<div class="btn" ng-click="leaveVoiceMessage({number: number, message: value})"></div>' +
		'',
		link: function(scope) {
			scope.networks = ["Verizon", "AT&T", "Pacific Bell"];
			scope.network = scope.networks[0];
		}
	}
})
