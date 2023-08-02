#include <queue>
#include <vector>

#include <iostream>
#include <algorithm>

#include <memory>

struct NoLimitTree final
{
public:
    struct Node final
    {
        // can add id.

        int value;
        std::vector<std::shared_ptr<Node>> lNodes;
        std::vector<std::shared_ptr<Node>> rNodes;

        explicit Node(int val) : value{ val } {}
    };

public:
    explicit NoLimitTree(size_t childCount = 2)
        : m_childCount{ childCount }, m_root{ nullptr } {}

    void clear() { m_root.reset(); }

public:
    size_t nodeCount() const {
        if (!m_root) return 0;

        size_t nodeCount{ 0 };
        std::queue<std::shared_ptr<Node>> qq;
        qq.push(m_root);

        while (!qq.empty()) {
            nodeCount++;

            const auto& top = qq.front();
            qq.pop();

            // ***

            for (const auto& lNode : top->lNodes)
                qq.push(lNode);

            for (const auto& rNode : top->rNodes)
                qq.push(rNode);
        }

        return nodeCount;
    }

public:
    void addNode(int val) {
        if (!m_root) {
            m_root = std::make_shared<Node>(val);
            return;
        }

        // ***

        std::queue<std::shared_ptr<Node>> qq;
        qq.push(m_root);

        while (!qq.empty()) {

            const auto top = qq.front();
            qq.pop();

            if (val < top->value) {
                if (top->lNodes.size() < m_childCount) {
                    top->lNodes.push_back(std::make_shared<Node>(val));
                    return;
                }
                else
                    for (const auto& lNode : top->lNodes)
                        qq.push(lNode);
            }
            else {
                if (top->rNodes.size() < m_childCount) {
                    top->rNodes.push_back(std::make_shared<Node>(val));
                    return;
                }
                else
                    for (const auto& rNode : top->rNodes)
                        qq.push(rNode);
            }
        }
    }

    void addNodes(std::vector<int> vals) {
        for (const auto& val : vals)
            addNode(val);
    }

public:
    void printTo(std::ostream& out) const {
        if (!m_root) {
            out << "<empty>;";
            return;
        }

        // ***

        std::queue<std::shared_ptr<Node>> qq;
        qq.push(m_root);

        while (!qq.empty()) {
            const auto top = qq.front();
            out << "    t: [ " << top->value << " ]\n";
            qq.pop();

            // ***

            out << "l: ";
            if (top->lNodes.empty())
                out << "<empty>";
            else
                for (const auto& lNode : top->lNodes) {
                    out << "[ " << lNode->value << " ] ";
                    qq.push(lNode);
                }
            out << std::endl;

            // ***

            out << "r: ";
            if (top->rNodes.empty())
                out << "<empty>";
            else
                for (const auto& rNode : top->rNodes) {
                    out << "[ " << rNode->value << " ] ";
                    qq.push(rNode);
                }
            out << std::endl;
        }
    }

private:
    const size_t m_childCount;
    std::shared_ptr<Node> m_root;
};

int main()
{
    NoLimitTree nlTree;
    std::cout << "Tree, node count: " << nlTree.nodeCount() << std::endl;

    nlTree.addNodes({ 1, 2, 3, 4, 5, -1, -2, -3, -4, -5 });
    std::cout << "Tree, node count: " << nlTree.nodeCount() << std::endl;

    nlTree.printTo(std::cout);

    // ***

    nlTree.clear();
    std::cout << "Tree, node count: " << nlTree.nodeCount() << std::endl;

    return 0;
}
