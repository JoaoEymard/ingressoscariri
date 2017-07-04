angular.module("ingressosCariri").controller('eventoCtrl', funcCtrl)

function funcCtrl($scope,$routeParams,$rootScope){
  $rootScope.pos_titulo = ' - '+$routeParams.titulo
  $scope.teste = $routeParams.titulo
  
  $scope.evento = {
    titulo: "Uma festa qualquer",
    img: "evento.png",
    data: "02/07/2017",
    hora: "22:00",
    cidade: "Juazeiro do Norte",
    estado: "CE",
    local: "Parque da Cidade",
    taxa: 0.15,
    periodo: [{
      data: "02/07/2017",
      hora: "22:00",
      atracao: "Jorge e Matheus, Safadão",
      lote: 1,
      categoria: [{
        nome: "VIP - Inteira",
        valor: 60.00,
        quantidade: 100
      },{
        nome: "VIP - Meia",
        valor: 30.00,
        quantidade: 50
      }],
    }],
  }
}
