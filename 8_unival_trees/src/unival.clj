;; Count unival subtrees of a tree

(ns unival
  "FIXME: my new org.corfield.new/scratch project.")

(deftype Node [^int val left right]
  Object
  (toString [this]
    (if this (str "Node(" val " Left=>" (if left (.toString left) "nil") " Right=>" (if right (.toString right) "nil") ")") "nil")))

;; I guess this could be rewritten to be much simpler and clearer, but for that I need to learn more Clojure
(defn isUnival? [tree]
  (cond
    (= tree nil) false
    (and (.left tree) (.right tree)) (and (isUnival? (.left tree)) (isUnival? (.right tree)) (= (.val tree) (->> tree .left .val) (->> tree .right .val)))
    (.left tree) (and (isUnival? (.left tree)) (= (.val tree) (->> tree .left .val)))
    (.right tree) (and (isUnival? (.right tree)) (= (.val tree) (->> tree .right .val)))
    :else true))

(defn traverseTree [tree f]
  (when tree
    (f tree)
    (traverseTree (.left tree) f)
    (traverseTree (.right tree) f)))

(defn countUnivalTrees [root]
  (let [cnt (atom 0)
        f (fn [tree]
            (when (isUnival? tree) (swap! cnt inc)))]
    (traverseTree root f)
    @cnt))

(defn countUnivalTreesBuiltin [root]
  (->> root
       (tree-seq #(not= % nil) #(vector (.left %1) (.right %1)))
       (filter isUnival?)
       (count)))

(def t1 (->Node 5 (->Node 5 nil nil) (->Node 5 nil nil)))
(->> t1 .right .val)
(isUnival? t1)
(isUnival? (->Node 5 nil nil))

(defn -main
  "Invoke me with clojure -M -m unival"
  [& args]
  (let [node (binding [*ns* (find-ns 'unival)] (load-string (apply str args)))]
    ;; (println "-main with" node (type node))
    ;; this workds also, you need to bind *ns* var to the namespace that you need,
    ;; bcs lost-string's default *ns* value is clojure.core
    (println (str node))
    (println "Is root tree unival? =>" (isUnival? node))
    (println "Number of unival subtrees =>" (countUnivalTreesBuiltin node))
    (println "Number of unival subtrees using builtin funcs =>" (countUnivalTrees node))))
;; so this doesn't work, there is no Node type in default load-string's context, i.e. clojure.core
;; (println (isUnival? (load-string node)))))

(def teststr "(->Node 5 (->Node 4 nil nil) nil)")
(isUnival? (load-string "(->Node 5 (->Node 4 nil nil) nil)"))
(isUnival? (load-string teststr))
