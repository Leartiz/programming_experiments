import QtQuick 2.15
import QtQuick.Window 2.15

import MyClass 1.0

Window {
    id: root

    width: 480
    height: 800

    visible: true
    title: qsTr("n23")

    Text {
        id: _text
        height: 50
        anchors.centerIn: parent
    }

    Timer {
        id: _timer
        interval: 1000
        repeat: false
        onTriggered: {
            Qt.quit()
        }
    }

    MyClass {
        id: _myClass

        onInitFailed: (msg) => {
            console.log(msg)

            _text.text = "init failed"

            root.color = "red"
            _timer.start()
        }

        Component.onCompleted: {
            _myClass.init()
        }
    }
}
